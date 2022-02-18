package action

/**
* Package used to register and list vehicles
* @package action
* @author Jordan Duarte
**/

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "github.com/jordanfduarte/vehicle-tracking-system/application"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
    "encoding/json"
)

type VehicleHandler struct {}

func NewVehicleHandler() *VehicleHandler {
   return &VehicleHandler{}
}

func (bc VehicleHandler) VehiclesGetAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    vehicles, err := application.GetAllVehicles()
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    JSON(w, http.StatusOK, vehicles)
}

func (bc VehicleHandler) VehiclesPostAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var vehicle *domain.Vehicles
    err2 := json.NewDecoder(r.Body).Decode(&vehicle)

    if err2 != nil {
        Error(w, http.StatusBadRequest, err2, "Invalid parameters entered")
        return
    }

    isValid, _ := vehicle.IsValid()
    if isValid == false {
        Error(w, http.StatusBadRequest, nil, "")
        return
    }

    fleet, err := application.GetRowFleet(vehicle.Fleet_ID)
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    if fleet.Fleet_ID <= 0 {
        Error(w, http.StatusNotFound, nil, "")
        return
    }

    err3 := application.AddVehicle(vehicle)
    if err3 != nil {
        Error(w, http.StatusNotFound, err3, err3.Error())
        return
    }

    returnTypeDynamic := &ReturnDynamic{Id: vehicle.Vehicle_ID}
    JSON(w, http.StatusCreated, returnTypeDynamic)
}
