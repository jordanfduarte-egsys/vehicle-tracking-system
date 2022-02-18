package repository

import (
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
)

/**
* Repository
* @package repository
* @author Jordan Duarte
**/

type VehicleCheckNullParam struct {
	Vehicles   *domain.Vehicles
	IsNullMaxSpeed bool
}

type VehiclesRepository interface {
    Get(id int) (*domain.Vehicles, error)
    GetAll() ([]domain.Vehicles, error)
    GetAllFleetAlertsByVehicle(id int) ([]domain.FleetAlerts, error)
    Save(*VehicleCheckNullParam) error
    RemoveAll() error
}