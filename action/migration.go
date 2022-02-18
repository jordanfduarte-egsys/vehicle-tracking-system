package action

/**
* Package used for initial migration of data to the bank
* @package action
* @author Jordan Duarte
**/

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "github.com/jordanfduarte/vehicle-tracking-system/config"
    "github.com/jordanfduarte/vehicle-tracking-system/application"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
)

type MigrationHandler struct {}

func NewMigrationHandler() *MigrationHandler {
   return &MigrationHandler{}
}

func (bc MigrationHandler) MigrationAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    _, err := config.DBMigrate()
    if err != nil {
        Error(w, http.StatusNotFound, err, err.Error())
        return
    }

    // Migrate Fleets initial
    application.AddFleet(&domain.Fleets{Name: "Veículos de perseguição", Max_Speed: 30.55});
    application.AddFleet(&domain.Fleets{Name: "Veículos de transporte de prisioneiros", Max_Speed: 25});
    application.AddFleet(&domain.Fleets{Name: "Escolta armada", Max_Speed: 22.22});

    log.Println("Migration has been processed")
    JSON(w, http.StatusOK, "Migration has been processed!")
}