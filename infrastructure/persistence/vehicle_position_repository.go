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

type VehiclePositionRepositoryImpl struct {
    Conn *gorm.DB
}

func VehiclePositionRepositoryWithRDB(conn *gorm.DB) repository.VehiclePositionsRepository {
    return &VehiclePositionRepositoryImpl{Conn: conn}
}

func (r *VehiclePositionRepositoryImpl) RemoveAll() error {
    return r.Conn.Exec( "DELETE FROM vehicle_positions" ).Error
}

func (r *VehiclePositionRepositoryImpl) GetAllByVeiches(id int) ([]domain.VehiclePositions, error) {
    vehiclePositions := []domain.VehiclePositions{}
    if err := r.Conn.Find(&vehiclePositions, "Vehicle_ID = ?", id).Error; err != nil {
        return nil, err
    }

    return vehiclePositions, nil
}

func (r *VehiclePositionRepositoryImpl) GetAll() ([]domain.VehiclePositions, error) {
    vehiclePositions := []domain.VehiclePositions{}
    if err := r.Conn.Preload("VehiclePosition").Find(&vehiclePositions).Error; err != nil {
        return nil, err
    }

    return vehiclePositions, nil
}

func (r *VehiclePositionRepositoryImpl) Save(vehicle *domain.VehiclePositions) error {
    if err := r.Conn.Create(&vehicle).Error; err != nil {
        return err
    }

    return nil
}
