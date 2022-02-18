package action

/**
* Package used to perform general data cleaning in mysql database
* @package action
* @author Jordan Duarte
**/

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "github.com/jordanfduarte/vehicle-tracking-system/application"
)

type DatabaseHandler struct {}

func NewDatabaseHandler() *DatabaseHandler {
   return &DatabaseHandler{}
}

func (bc DatabaseHandler) DatabaseAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    err := application.RemoveVehiclePositionAll()
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    err2 := application.RemoveVehicleAll()
    if err2 != nil {
        Error(w, http.StatusNotFound, err2, err2.Error())
        return
    }

    err3 := application.RemoveFleetAlertAll()
    if err3 != nil {
        Error(w, http.StatusNotFound, err3, err3.Error())
        return
    }

    err4 := application.RemoveFleetAll()
    if err4 != nil {
        Error(w, http.StatusNotFound, err4, err4.Error())
        return
    }

    JSON(w, http.StatusOK, "Delete has been processed")
}