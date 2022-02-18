package action

/**
* Package used to register and list a fleet
* @package action
* @author Jordan Duarte
**/

import (
    "encoding/json"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "github.com/jordanfduarte/vehicle-tracking-system/application"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
)

type FleetHandler struct {}

func NewFleetHandler() *FleetHandler {
   return &FleetHandler{}
}

func (bc FleetHandler) FleetsGetAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fleets, err := application.GetAllFleets()

    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    JSON(w, http.StatusOK, fleets)
}

func (bc FleetHandler) FleetsPostAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var fleet *domain.Fleets
    err := json.NewDecoder(r.Body).Decode(&fleet)

    if err != nil {
        Error(w, http.StatusBadRequest, err, "Invalid parameters entered")
        return
    }

    isValid, error := fleet.IsValid()
    if isValid == false {
        Error(w, http.StatusBadRequest, nil, error)
        return
    }

    err2 := application.AddFleet(fleet)
    if err2 != nil {
        Error(w, http.StatusNotFound, err2, err2.Error())
        return
    }

    ReturnDynamic := &ReturnDynamic{Id: fleet.Fleet_ID}
    JSON(w, http.StatusCreated, ReturnDynamic)
}