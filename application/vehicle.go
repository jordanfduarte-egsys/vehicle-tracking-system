package application

import (
	//"github.com/biezhi/gorm-paginator/pagination"

	"../config"
	"../domain"
	"../infrastructure/persistence"
)

func RemoveVehicleAll() error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.VehicleRepositoryWithRDB(conn)
	return repo.RemoveAll()
}
