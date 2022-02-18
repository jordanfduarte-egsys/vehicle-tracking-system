package action

/**
* Package used to register and return vehicle positions
* @package action
* @author Jordan Duarte
**/

import (
    "encoding/json"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "strconv"
    "github.com/jordanfduarte/vehicle-tracking-system/application"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
)

type VehiclePositionHandler struct {}

func NewVehiclePositionHandler() *VehiclePositionHandler {
   return &VehiclePositionHandler{}
}

func (bc VehiclePositionHandler) PositionsGetAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    param := ps.ByName("id")
    id, err := strconv.Atoi(param)
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    positions, err := application.GetAllPositionsByVehicles(id)
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    JSON(w, http.StatusOK, positions)
}

func (bc VehiclePositionHandler) PositionsPostAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    param := ps.ByName("id")
    id, err := strconv.Atoi(param)
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    var vehiclePosition *domain.VehiclePositions
    err2 := json.NewDecoder(r.Body).Decode(&vehiclePosition)

    if err2 != nil {
        Error(w, http.StatusBadRequest, err2, "Invalid parameters entered")
        return
    }

    vehicle, err := application.GetRowVehicle(id)
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    if vehicle.Vehicle_ID <= 0 {
        Error(w, http.StatusNotFound, nil, "")
        return
    }

    isValid, _:= vehiclePosition.IsValid()

    if isValid == false {
        Error(w, http.StatusBadRequest, nil, "")
        return
    }

    vehiclePosition.Vehicle_ID = id
    vehicleFind, err65 := application.GetRowVehicle(id)

    if err65 != nil {
        Error(w, http.StatusNotFound, err65, err65.Error())
        return
    }

    vehiclePosition.Max_Speed = vehicleFind.Max_Speed
    err3 := application.AddPositionVehicle(vehiclePosition)
    if err3 != nil {
        Error(w, http.StatusNotFound, err3, err3.Error())
        return
    }

    if vehiclePosition.Current_Speed > vehiclePosition.Max_Speed {
        fleetAlerts, err5 := application.GetAllFleetAlertsByVehicle(id)
        if err5 != nil {
            Error(w, http.StatusNotFound, err5, err5.Error())
            return
        }

        // Worker Pools
        jobs := make(chan domain.Site, 100)
        for w := 1; w <= 3; w++ {
            go domain.Crawl(w, jobs)
        }

        for _, fleetAlertSend := range fleetAlerts {
            jsonValue, _ := json.Marshal(vehiclePosition)
            jobs <- domain.Site{URL: fleetAlertSend.WebHook, Buffer: jsonValue}
        }
        close(jobs)
    }

    returnTypeDynamic := &ReturnDynamic{Id: vehiclePosition.Vehicle_Position_ID}
    JSON(w, http.StatusCreated, returnTypeDynamic)
}