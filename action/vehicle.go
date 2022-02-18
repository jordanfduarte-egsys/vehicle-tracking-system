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
    "github.com/jordanfduarte/vehicle-tracking-system/domain/repository"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
    "encoding/json"
    // "log"
    "strconv"
    "io"
    "fmt"
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
    vehicle := &domain.Vehicles{}

    dec := json.NewDecoder(r.Body)
    var validNext = false
    var nameInit string
    IsNullMaxSpeed := true
    for {
        t, err := dec.Token()
        if err == io.EOF {
            break
        }
        if err != nil {
            continue
        }
        //log.Printf("%T: %v", t, t)
        value := fmt.Sprintf("%v", t)

        if value == "fleet_id" {
            validNext = true
            nameInit = "fleet_id"
            continue
        }

        if value == "name" && !validNext {
            validNext = true
            nameInit = "name"
            continue
        }

        if value == "max_speed" && !validNext {
            validNext = true
            nameInit = "max_speed"
            continue
        }

        if validNext {
            if nameInit == "fleet_id" {
                idFleet, err := strconv.Atoi(fmt.Sprintf("%v", t))
                if err != nil {
                    continue
                }
                vehicle.Fleet_ID = idFleet
                validNext = false
            }

            if nameInit == "name" {
                vehicle.Name = fmt.Sprintf("%v", t)
                validNext = false
            }

            if nameInit == "max_speed" {
                value, err := strconv.ParseFloat(fmt.Sprintf("%v", t), 32)
                if err != nil {
                    continue
                }
                vehicle.Max_Speed = float32(value)
                validNext = false
                IsNullMaxSpeed = false
            }
        }
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

    vehicleCheckNullParam := &repository.VehicleCheckNullParam{
        Vehicles: vehicle,
        IsNullMaxSpeed: IsNullMaxSpeed}
    err3 := application.AddVehicle(vehicleCheckNullParam)
    if err3 != nil {
        Error(w, http.StatusNotFound, err3, err3.Error())
        return
    }

    returnTypeDynamic := &ReturnDynamic{Id: vehicle.Vehicle_ID}
    JSON(w, http.StatusCreated, returnTypeDynamic)
}
