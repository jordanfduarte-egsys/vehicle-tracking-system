package application

import (
	//"github.com/biezhi/gorm-paginator/pagination"

	"github.com/jordanfduarte/vehicle-tracking-system/config"
	"github.com/jordanfduarte/vehicle-tracking-system/domain"
	"github.com/jordanfduarte/vehicle-tracking-system/infrastructure/persistence"
)

func RemoveVehicleAll() error {
	conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
	if err != nil {
		return err
	}
	// defer conn.Close()

	repo := persistence.VehicleRepositoryWithRDB(conn)
	// vehicles, err := repo.GetAll()
	// return repo.RemoveAll(vehicles)
	return repo.RemoveAll()
}

func AddVehicle(vehicle *domain.Vehicles) error {
	conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
	if err != nil {
		return err
	}
	// defer conn.Close()

	repo := persistence.VehicleRepositoryWithRDB(conn)
	return repo.Save(vehicle)
}


func GetAllVehicles() ([]domain.Vehicles, error) {
	conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	repo := persistence.VehicleRepositoryWithRDB(conn)
	// vehicles, err := repo.GetAll()
	// return repo.RemoveAll(vehicles)
	return repo.GetAll()
}


func GetRowVehicle(id int) (*domain.Vehicles, error) {
	conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	repo := persistence.VehicleRepositoryWithRDB(conn)
	// vehicles, err := repo.GetAll()
	// return repo.RemoveAll(vehicles)
	return repo.Get(id)
}

func GetAllFleetAlertsByVehicle(id int) ([]domain.FleetAlerts, error) {
	conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	repo := persistence.VehicleRepositoryWithRDB(conn)
	// vehicles, err := repo.GetAll()
	// return repo.RemoveAll(vehicles)
	return repo.GetAllFleetAlertsByVehicle(id)
}
