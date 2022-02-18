package application

/**
* Package used for fleet business logic
* @package application
* @author Jordan Duarte
**/

import (
    "github.com/jordanfduarte/vehicle-tracking-system/config"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
    "github.com/jordanfduarte/vehicle-tracking-system/infrastructure/persistence"
)

func AddFleet(fleet *domain.Fleets) error {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return err
    }
    repo := persistence.FleetRepositoryWithRDB(conn)
    return repo.Save(fleet)
}

func RemoveFleetAll() error {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return err
    }
    repo := persistence.FleetRepositoryWithRDB(conn)
    return repo.RemoveAll()
}

func GetAllFleets() ([]domain.Fleets, error) {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return nil, err
    }

    repo := persistence.FleetRepositoryWithRDB(conn)
    return repo.GetAll()
}

func GetRowFleet(id int) (*domain.Fleets, error) {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return nil, err
    }

    repo := persistence.FleetRepositoryWithRDB(conn)
    return repo.Get(id)
}

func GetFirstFleetRow() (*domain.Fleets, error) {
    conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
    if err != nil {
        return nil, err
    }

    repo := persistence.FleetRepositoryWithRDB(conn)
    return repo.GetFirst()
}

