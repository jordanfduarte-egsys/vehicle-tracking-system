package config

import (
	"log"

	"github.com/jinzhu/gorm"
	"../domain"
)

// DBMigrate will create & migrate the tables, then make the some relationships if necessary
func DBMigrate() (*gorm.DB, error) {
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	conn.AutoMigrate(domain.Fleet_Alert{}, domain.Fleet{}, domain.Vehicle_Position{}, domain.Vehicle{})
	log.Println("Migration has been processed")

	return conn, nil
}