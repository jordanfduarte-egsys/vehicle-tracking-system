package repository

import "../../domain"

// VehiclePosition represent repository of the Vehicle Position
// Expect implementation by the infrastructure layer
type VehiclePositionRepository interface {
	Get(id int) (*domain.VehiclePosition, error)
	GetAll() ([]domain.VehiclePosition, error)
	Save(*domain.VehiclePosition) error
	Remove(id int) error
	RemoveAll() error
	Update(*domain.VehiclePosition) error
}