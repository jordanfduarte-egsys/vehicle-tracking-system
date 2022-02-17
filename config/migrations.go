package config

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"github.com/jordanfduarte/vehicle-tracking-system/domain"
	//"github.com/jordanfduarte/vehicle-tracking-system/application"
)

// DBMigrate will create & migrate the tables, then make the some relationships if necessary
func DBMigrate() (*gorm.DB, error) {
	conn, err := ConnectDB(&Options{IsDefaultDbName: false})
	// var config = ConfigDB{}
	// config.Read()

	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	queryCreateDatabase := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.Dbname)
	useDatatable := fmt.Sprintf("USE %s", config.Dbname)
	log.Printf("[DEBUG] "+queryCreateDatabase)
	conn.Exec(queryCreateDatabase)
	conn.Exec(useDatatable)
	conn.AutoMigrate(domain.FleetAlerts{}, domain.Fleets{}, domain.VehiclePositions{}, domain.Vehicles{})

	return conn, nil
}