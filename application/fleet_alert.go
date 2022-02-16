package application

import (
	//"github.com/biezhi/gorm-paginator/pagination"

	"../config"
	"../domain"
	"../infrastructure/persistence"
)

// RemoveNews do remove news by id
func RemoveFleetAlertAll() error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.RemoveAll()
}