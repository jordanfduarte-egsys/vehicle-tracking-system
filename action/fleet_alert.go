package action

/**
* Package used to register a fleet notification
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

type AlertsHandler struct {}

func NewAlertsHandler() *AlertsHandler {
   return &AlertsHandler{}
}

func (bc AlertsHandler) AlertsGetAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    param := ps.ByName("id")
    id, err := strconv.Atoi(param)
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    // veririca se exist o Fleet
    fleet, err := application.GetRowFleet(id)
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    if fleet.Fleet_ID <= 0 {
        Error(w, http.StatusNotFound, nil, "")
        return
    }

    fleetAlerts, err := application.GetAllFleetAlerts(id)
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    JSON(w, http.StatusOK, fleetAlerts)
}

func (bc AlertsHandler) AlertsPostAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    param := ps.ByName("id")
    id, err := strconv.Atoi(param)
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    var fleetAlert *domain.FleetAlerts
    err2 := json.NewDecoder(r.Body).Decode(&fleetAlert)

    if err2 != nil {
        Error(w, http.StatusBadRequest, err2, "Invalid parameters entered")
        return
    }

    fleet, err := application.GetRowFleet(id)
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    if fleet.Fleet_ID <= 0 {
        Error(w, http.StatusNotFound, nil, "")
        return
    }

    isValid, _ := fleetAlert.IsValid()
    if isValid == false {
        Error(w, http.StatusBadRequest, nil, "")
        return
    }

    fleetAlert.Fleet_ID = id
    err3 := application.AddFleetAlert(fleetAlert)
    if err3 != nil {
        Error(w, http.StatusNotFound, err3, err3.Error())
        return
    }

    returnTypeDynamic := &ReturnDynamic{Id: fleetAlert.Fleet_Alert_ID}
    JSON(w, http.StatusCreated, returnTypeDynamic)
}
