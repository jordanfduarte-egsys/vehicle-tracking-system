package repository

import "github.com/jordanfduarte/vehicle-tracking-system/domain"

/**
* Repository
* @package repository
* @author Jordan Duarte
**/

type VehiclesRepository interface {
    Get(id int) (*domain.Vehicles, error)
    GetAll() ([]domain.Vehicles, error)
    GetAllFleetAlertsByVehicle(id int) ([]domain.FleetAlerts, error)
    Save(*domain.Vehicles) error
    RemoveAll() error
}