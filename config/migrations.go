package config

/**
* Package used to configure migrations
* @package config
* @author Jordan Duarte
**/

import (
    "fmt"
    "log"
    "gorm.io/gorm"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
)

// DBMigrate will create & migrate the tables, then make the some relationships if necessary
func DBMigrate() (*gorm.DB, error) {
    conn, err := ConnectDB(&Options{IsDefaultDbName: false})

    if err != nil {
        return nil, err
    }

    queryCreateDatabase := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.Dbname)
    useDatatable := fmt.Sprintf("USE %s", config.Dbname)
    log.Printf("[DEBUG] "+queryCreateDatabase)
    conn.Exec(queryCreateDatabase)
    conn.Exec(useDatatable)
    conn.AutoMigrate(domain.FleetAlerts{}, domain.Fleets{}, domain.VehiclePositions{}, domain.Vehicles{})

    return conn, nil
}