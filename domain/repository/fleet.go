package repository

import "../../domain"

// FleetRepository represent repository of the fleet
// Expect implementation by the infrastructure layer
type FleetRepository interface {
	Get(id int) (*domain.Fleet, error)
	GetAll() ([]domain.Fleet, error)
	Save(*domain.Fleet) error
	Remove(id int) error
	RemoveAll() error
	Update(*domain.Fleet) error
}