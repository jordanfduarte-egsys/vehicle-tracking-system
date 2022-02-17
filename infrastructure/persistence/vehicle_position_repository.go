package persistence

import (
	"gorm.io/gorm"
	//_ "gorm.io/gorm/dialects/mysql"
	"github.com/jordanfduarte/vehicle-tracking-system/domain"
	"github.com/jordanfduarte/vehicle-tracking-system/domain/repository"
)

// VehiclePositionRepositoryImpl Implements repository.VehiclePositionRepository
type VehiclePositionRepositoryImpl struct {
	Conn *gorm.DB
}

// VehiclePositionRepositoryWithRDB returns initialized VehiclePositionRepositoryImpl
func VehiclePositionRepositoryWithRDB(conn *gorm.DB) repository.VehiclePositionsRepository {
	return &VehiclePositionRepositoryImpl{Conn: conn}
}

func (r *VehiclePositionRepositoryImpl) Remove(id int) error {
	return nil
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

func (r *VehiclePositionRepositoryImpl) Get(id int) (*domain.VehiclePositions, error) {
	return nil, nil
}

func (r *VehiclePositionRepositoryImpl) Save(vehicle *domain.VehiclePositions) error {
	if err := r.Conn.Create(&vehicle).Error; err != nil {
		return err
	}

	return nil
}

func (r *VehiclePositionRepositoryImpl) Update(*domain.VehiclePositions) error {
	return nil
}