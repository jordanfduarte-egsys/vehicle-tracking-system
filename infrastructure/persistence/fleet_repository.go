package persistence

import (
	"gorm.io/gorm"
	//_ "gorm.io/gorm/dialects/mysql"
	"github.com/jordanfduarte/vehicle-tracking-system/domain"
	"github.com/jordanfduarte/vehicle-tracking-system/domain/repository"
)

// FleetRepositoryImpl Implements repository.FleetsRepositoryImpl
type FleetRepositoryImpl struct {
	Conn *gorm.DB
}

// FleetRepositoryImplWithRDB returns initialized FleetRepositoryImpl
func FleetRepositoryWithRDB(conn *gorm.DB) repository.FleetsRepository {
	return &FleetRepositoryImpl{Conn: conn}
}

func (r *FleetRepositoryImpl) Remove(id int) error {
	return nil
}

func (r *FleetRepositoryImpl) RemoveAll() error {
	return r.Conn.Exec( "DELETE FROM fleets" ).Error
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

func (r *FleetRepositoryImpl) Update(*domain.Fleets) error {
	return nil
}