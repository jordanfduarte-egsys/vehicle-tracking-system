package repository

import "github.com/jordanfduarte/vehicle-tracking-system/domain"

/**
* Repository
* @package repository
* @author Jordan Duarte
**/

type VehiclePositionsRepository interface {
    GetAll() ([]domain.VehiclePositions, error)
    GetAllByVeiches(id int) ([]domain.VehiclePositions, error)
    Save(*domain.VehiclePositions) error
    RemoveAll() error
}