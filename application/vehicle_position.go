package application

import (
	//"github.com/biezhi/gorm-paginator/pagination"

	"github.com/jordanfduarte/vehicle-tracking-system/config"
	"github.com/jordanfduarte/vehicle-tracking-system/domain"
	"github.com/jordanfduarte/vehicle-tracking-system/infrastructure/persistence"
)


// RemoveNews do remove news by id
func RemoveVehiclePositionAll() error {
	conn, err := config.ConnectDB(&config.Options{IsDefaultDbName: true})
	if err != nil {
		return err
	}
	// defer conn.Close()

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
	// defer conn.Close()

	repo := persistence.VehiclePositionRepositoryWithRDB(conn)
	return repo.Save(vehiclePosition)
}