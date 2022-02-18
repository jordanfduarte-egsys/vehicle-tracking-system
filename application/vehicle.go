package application

/**
* Package used for vehicle business logic
* @package application
* @author Jordan Duarte
**/

import (
    "github.com/jordanfduarte/vehicle-tracking-system/config"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
    "github.com/jordanfduarte/vehicle-tracking-system/infrastructure/persistence"
    "github.com/jordanfduarte/vehicle-tracking-system/domain/repository"
)

func RemoveVehicleAll() error {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return err
    }

    repo := persistence.VehicleRepositoryWithRDB(conn)
    return repo.RemoveAll()
}

func AddVehicle(vehicle *repository.VehicleCheckNullParam) error {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return err
    }

    repo := persistence.VehicleRepositoryWithRDB(conn)
    return repo.Save(vehicle)
}


func GetAllVehicles() ([]domain.Vehicles, error) {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return nil, err
    }

    repo := persistence.VehicleRepositoryWithRDB(conn)
    return repo.GetAll()
}


func GetRowVehicle(id int) (*domain.Vehicles, error) {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return nil, err
    }

    repo := persistence.VehicleRepositoryWithRDB(conn)
    return repo.Get(id)
}

func GetAllFleetAlertsByVehicle(id int) ([]domain.FleetAlerts, error) {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return nil, err
    }
    repo := persistence.VehicleRepositoryWithRDB(conn)
    return repo.GetAllFleetAlertsByVehicle(id)
}
