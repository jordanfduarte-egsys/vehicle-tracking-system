package repository

import "github.com/jordanfduarte/vehicle-tracking-system/domain"

// Vehicle represent repository of the Vehicle
// Expect implementation by the infrastructure layer
type VehiclesRepository interface {
	Get(id int) (*domain.Vehicles, error)
	GetAll() ([]domain.Vehicles, error)
	GetAllFleetAlertsByVehicle(id int) ([]domain.FleetAlerts, error)
	Save(*domain.Vehicles) error
	Remove(id int) error
	RemoveAll() error
	Update(*domain.Vehicles) error
}