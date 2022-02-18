package application

/**
* Package used for vehicle position business logic
* @package application
* @author Jordan Duarte
**/

import (
    "github.com/jordanfduarte/vehicle-tracking-system/config"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
    "github.com/jordanfduarte/vehicle-tracking-system/infrastructure/persistence"
)

func RemoveVehiclePositionAll() error {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return err
    }
    repo := persistence.VehiclePositionRepositoryWithRDB(conn)
    return repo.RemoveAll()
}

func GetAllPositionsByVehicles(id int) ([]domain.VehiclePositions, error) {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return nil, err
    }

    repo := persistence.VehiclePositionRepositoryWithRDB(conn)
    return repo.GetAllByVeiches(id)
}

func AddPositionVehicle(vehiclePosition *domain.VehiclePositions) error {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return err
    }
    repo := persistence.VehiclePositionRepositoryWithRDB(conn)
    return repo.Save(vehiclePosition)
}