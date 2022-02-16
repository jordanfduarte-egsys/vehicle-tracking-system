package repository

import "../../domain"

// FleetAlertRepository represent repository of  the fleet alert
// Expect implementation by the infrastructure layer
type FleetAlertRepository interface {
	Get(id int) (*domain.Fleet_Alert, error)
	GetAll() ([]domain.Fleet_Alert, error)
	Save(*domain.Fleet_Alert) error
	Remove(id int) error
	RemoveAll() error
	Update(*domain.Fleet_Alert) error
}