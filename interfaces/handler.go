package interfaces

/**
* HTTP handler initial
* @package interfaces
* @author Jordan Duarte
**/

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "github.com/jordanfduarte/vehicle-tracking-system/action"
)

// Run start server
func Run(port int) error {
    log.Printf("Server running at http://localhost:%d/", port)
    return http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
    r := httprouter.New()

    databaseHandler := action.NewDatabaseHandler()
    fleetAlertHandler := action.NewAlertsHandler()
    fleetHandler := action.NewFleetHandler()
    indexHandler := action.NewIndexHandler()
    migrationHandler := action.NewMigrationHandler()
    vehiclePositionHandler := action.NewVehiclePositionHandler()
    vehicleHandler := action.NewVehicleHandler()

    /**
    * Init Route
    */
    r.GET("/", indexHandler.IndexAction)

    /**
    * Reset
    */
    // limpa toda a base de dados
    r.DELETE("/database", databaseHandler.DatabaseAction)

    /**
    * Fleets
    */
    // lista todas as frotas
    r.GET("/fleets", fleetHandler.FleetsGetAction)

    // Cria uma frota
    r.POST("/fleets", fleetHandler.FleetsPostAction)

    // lista todas os alertas de uma frota
    r.GET("/fleets/:id/alerts", fleetAlertHandler.AlertsGetAction)

    // Cria uma alerta para frota
    r.POST("/fleets/:id/alerts", fleetAlertHandler.AlertsPostAction)

    /**
    * Vehicles
    */
    // lista todos os veículos
    r.GET("/vehicles", vehicleHandler.VehiclesGetAction)

    // Cria uma veículo
    r.POST("/vehicles", vehicleHandler.VehiclesPostAction)

    /*
    * Vehicles Positions
    */
    // lista todos as posições de um veículo
    r.GET("/vehicles/:id/positions", vehiclePositionHandler.PositionsGetAction)

    // Salva a posição do veículo)
    r.POST("/vehicles/:id/positions", vehiclePositionHandler.PositionsPostAction)

    // Migrations
    r.GET("/migration", migrationHandler.MigrationAction)

    return r
}



