package application

/**
* Package used for fleet alert business logic
* @package application
* @author Jordan Duarte
**/

import (
    "github.com/jordanfduarte/vehicle-tracking-system/config"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
    "github.com/jordanfduarte/vehicle-tracking-system/infrastructure/persistence"
)

func RemoveFleetAlertAll() error {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return err
    }

    repo := persistence.FleetAlertRepositoryWithRDB(conn)
    return repo.RemoveAll()
}

func GetAllFleetAlerts(id int) ([]domain.FleetAlerts, error) {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return nil, err
    }

    repo := persistence.FleetAlertRepositoryWithRDB(conn)
    return repo.GetAll(id)
}

func AddFleetAlert(fleetAlert *domain.FleetAlerts) error {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return err
    }

    repo := persistence.FleetAlertRepositoryWithRDB(conn)
    return repo.Save(fleetAlert)
}
