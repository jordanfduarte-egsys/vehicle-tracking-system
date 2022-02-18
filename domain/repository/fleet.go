package repository

import "github.com/jordanfduarte/vehicle-tracking-system/domain"

/**
* Repository
* @package repository
* @author Jordan Duarte
**/

type FleetsRepository interface {
    Get(id int) (*domain.Fleets, error)
    GetAll() ([]domain.Fleets, error)
    Save(*domain.Fleets) error
    RemoveAll() error
}