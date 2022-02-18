package persistence

/**
* Implements repository interface
* @package persistence
* @author Jordan Duarte
**/

import (
    "gorm.io/gorm"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
    "github.com/jordanfduarte/vehicle-tracking-system/domain/repository"
)

type FleetAlertRepositoryImpl struct {
    Conn *gorm.DB
}

func FleetAlertRepositoryWithRDB(conn *gorm.DB) repository.FleetAlertsRepository {
    return &FleetAlertRepositoryImpl{Conn: conn}
}

func (r *FleetAlertRepositoryImpl) RemoveAll() error {
    e := r.Conn.Exec("DELETE FROM fleet_alerts").Error
    if e != nil {
        return e
    }
    return r.Conn.Exec("ALTER TABLE fleet_alerts AUTO_INCREMENT=0;").Error
}

func (r *FleetAlertRepositoryImpl) GetAll(id int) ([]domain.FleetAlerts, error) {
    fleetAlerts := []domain.FleetAlerts{}
    if err := r.Conn.Find(&fleetAlerts, "Fleet_ID = ?", id).Error; err != nil {
        return nil, err
    }

    return fleetAlerts, nil
}

func (r *FleetAlertRepositoryImpl) Save(fleetAlert *domain.FleetAlerts) error {
    if err := r.Conn.Create(&fleetAlert).Error; err != nil {
        return err
    }

    return nil
}
