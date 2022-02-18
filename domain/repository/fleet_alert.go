package repository

import "github.com/jordanfduarte/vehicle-tracking-system/domain"

/**
* Repository
* @package repository
* @author Jordan Duarte
**/

type FleetAlertsRepository interface {
    GetAll(id int) ([]domain.FleetAlerts, error)
    Save(*domain.FleetAlerts) error
    RemoveAll() error
}