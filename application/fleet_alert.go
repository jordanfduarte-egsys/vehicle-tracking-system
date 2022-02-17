package application

import (
	//"github.com/biezhi/gorm-paginator/pagination"

	"github.com/jordanfduarte/vehicle-tracking-system/config"
	"github.com/jordanfduarte/vehicle-tracking-system/domain"
	"github.com/jordanfduarte/vehicle-tracking-system/infrastructure/persistence"
)

// RemoveNews do remove news by id
func RemoveFleetAlertAll() error {
	conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
	if err != nil {
		return err
	}
	// defer conn.Close()

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
	// defer conn.Close()

	repo := persistence.FleetAlertRepositoryWithRDB(conn)
	return repo.Save(fleetAlert)
}
