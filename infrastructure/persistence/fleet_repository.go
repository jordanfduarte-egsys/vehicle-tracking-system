package persistence

/**
* Implements repository interface
* @package persistence
* @author Jordan Duarte
**/

import (
    "gorm.io/gorm"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
    "github.com/jordanfduarte/vehicle-tracking-system/domain/repository"
)

type FleetRepositoryImpl struct {
    Conn *gorm.DB
}

func FleetRepositoryWithRDB(conn *gorm.DB) repository.FleetsRepository {
    return &FleetRepositoryImpl{Conn: conn}
}

func (r *FleetRepositoryImpl) RemoveAll() error {
    e := r.Conn.Exec("DELETE FROM fleets").Error
    if e != nil {
        return e
    }
    return r.Conn.Exec("ALTER TABLE fleets AUTO_INCREMENT=0;").Error
}

func (r *FleetRepositoryImpl) GetFirst() (*domain.Fleets, error) {
    fleets := &domain.Fleets{}
    r.Conn.Raw(`
        SELECT * FROM fleets ORDER BY Fleet_ID ASC LIMIT 1`).Find(&fleets)

    return fleets, nil
}

func (r *FleetRepositoryImpl) GetAll() ([]domain.Fleets, error) {
    fleets := []domain.Fleets{}
    if err := r.Conn.Find(&fleets).Error; err != nil {
        return nil, err
    }

    return fleets, nil
}

func (r *FleetRepositoryImpl) Get(id int) (*domain.Fleets, error) {
    fleet := &domain.Fleets{}
    if err := r.Conn.Find(fleet, "Fleet_ID = ?", id).Error; err != nil {
        return nil, err
    }

    return fleet, nil
}

func (r *FleetRepositoryImpl) Save(fleet *domain.Fleets) error {
    if err := r.Conn.Create(&fleet).Error; err != nil {
        return err
    }

    return nil
}
