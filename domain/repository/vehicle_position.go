package repository

import "github.com/jordanfduarte/vehicle-tracking-system/domain"

// VehiclePosition represent repository of the Vehicle Position
// Expect implementation by the infrastructure layer
type VehiclePositionsRepository interface {
	Get(id int) (*domain.VehiclePositions, error)
	GetAll() ([]domain.VehiclePositions, error)
	GetAllByVeiches(id int) ([]domain.VehiclePositions, error)
	Save(*domain.VehiclePositions) error
	Remove(id int) error
	RemoveAll() error
	Update(*domain.VehiclePositions) error
}