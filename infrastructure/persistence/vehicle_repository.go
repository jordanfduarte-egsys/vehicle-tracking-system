package persistence

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"../../domain"
	"../../domain/repository"
)

type VehicleRepositoryImpl struct {
	Conn *gorm.DB
}

func VehicleRepositoryWithRDB(conn *gorm.DB) repository.VehicleRepository {
	return &VehicleRepositoryImpl{Conn: conn}
}

func (r *VehicleRepositoryImpl) RemoveAll() error {
	tx := r.Conn.Begin()
	if err := tx.Delete().Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}