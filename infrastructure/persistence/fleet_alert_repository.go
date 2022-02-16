package persistence

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"../../domain"
	"../../domain/repository"
)

// FleetAlertRepositoryImpl Implements repository.FleetAlertRepository
type FleetAlertRepositoryImpl struct {
	Conn *gorm.DB
}

// FleetAlertNewsRepositoryWithRDB returns initialized FleetAlertRepositoryImpl
func FleetAlertRepositoryWithRDB(conn *gorm.DB) repository.FleetAlertRepository {
	return &FleetAlertRepositoryImpl{Conn: conn}
}

func (r *FleetAlertRepositoryImpl) Remove(id int) error {
	tx := r.Conn.Begin()
	if err := tx.Delete().Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}