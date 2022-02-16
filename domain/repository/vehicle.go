package repository

import "../../domain"

// Vehicle represent repository of the Vehicle
// Expect implementation by the infrastructure layer
type VehicleRepository interface {
	Get(id int) (*domain.Vehicle, error)
	GetAll() ([]domain.Vehicle, error)
	Save(*domain.Vehicle) error
	Remove(id int) error
	RemoveAll() error
	Update(*domain.Vehicle) error
}