package repository

import "github.com/jordanfduarte/vehicle-tracking-system/domain"

// FleetAlertRepository represent repository of  the fleet alert
// Expect implementation by the infrastructure layer
type FleetAlertsRepository interface {
	Get(id int) (*domain.FleetAlerts, error)
	GetAll(id int) ([]domain.FleetAlerts, error)
	Save(*domain.FleetAlerts) error
	Remove(id int) error
	RemoveAll() error
	Update(*domain.FleetAlerts) error
}