package repository

import "github.com/jordanfduarte/vehicle-tracking-system/domain"

// FleetRepository represent repository of the fleet
// Expect implementation by the infrastructure layer
type FleetsRepository interface {
	Get(id int) (*domain.Fleets, error)
	GetAll() ([]domain.Fleets, error)
	Save(*domain.Fleets) error
	Remove(id int) error
	RemoveAll() error
	Update(*domain.Fleets) error
}