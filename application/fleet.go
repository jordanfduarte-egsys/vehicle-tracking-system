package application

import (
	//"github.com/biezhi/gorm-paginator/pagination"

	"github.com/jordanfduarte/vehicle-tracking-system/config"
	"github.com/jordanfduarte/vehicle-tracking-system/domain"
	"github.com/jordanfduarte/vehicle-tracking-system/infrastructure/persistence"
)

func AddFleet(fleet *domain.Fleets) error {
	conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
	if err != nil {
		return err
	}
	// defer conn.Close()

	repo := persistence.FleetRepositoryWithRDB(conn)
	return repo.Save(fleet)
}

// RemoveNews do remove news by id
func RemoveFleetAll() error {
	conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
	if err != nil {
		return err
	}
	// defer conn.Close()

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

// func AddFleet(fleet *domain.Fleets) (domain.Fleets, error) {
// 	conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
// 	if err != nil {
// 		return nil, err
// 	}

// 	repo := persistence.FleetRepositoryWithRDB(conn)
// 	return repo.Save(fleet)
// }
