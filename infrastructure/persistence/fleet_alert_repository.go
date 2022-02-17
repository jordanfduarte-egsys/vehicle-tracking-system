package persistence

import (
	"gorm.io/gorm"
	//_ "gorm.io/gorm/dialects/mysql"
	"github.com/jordanfduarte/vehicle-tracking-system/domain"
	"github.com/jordanfduarte/vehicle-tracking-system/domain/repository"
)

// FleetAlertRepositoryImpl Implements repository.FleetAlertsRepository
type FleetAlertRepositoryImpl struct {
	Conn *gorm.DB
}

// FleetAlertNewsRepositoryWithRDB returns initialized FleetAlertRepositoryImpl
func FleetAlertRepositoryWithRDB(conn *gorm.DB) repository.FleetAlertsRepository {
	return &FleetAlertRepositoryImpl{Conn: conn}
}

func (r *FleetAlertRepositoryImpl) Remove(id int) error {
	return nil
}

func (r *FleetAlertRepositoryImpl) RemoveAll() error {
	return r.Conn.Exec("DELETE FROM fleet_alerts").Error
}

func (r *FleetAlertRepositoryImpl) GetAll(id int) ([]domain.FleetAlerts, error) {
	fleetAlerts := []domain.FleetAlerts{}
	if err := r.Conn.Find(&fleetAlerts, "Fleet_ID = ?", id).Error; err != nil {
		return nil, err
	}

	return fleetAlerts, nil
}

func (r *FleetAlertRepositoryImpl) Get(id int) (*domain.FleetAlerts, error) {
	return nil, nil
}

func (r *FleetAlertRepositoryImpl) Save(fleetAlert *domain.FleetAlerts) error {
	if err := r.Conn.Create(&fleetAlert).Error; err != nil {
		return err
	}

	return nil
}

func (r *FleetAlertRepositoryImpl) Update(*domain.FleetAlerts) error {
	return nil
}