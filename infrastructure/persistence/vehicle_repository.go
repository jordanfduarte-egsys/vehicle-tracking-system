package persistence

import (
	"gorm.io/gorm"
	// _ "gorm.io/gorm/dialects/mysql"
	// "github.com/jordanfduarte/vehicle-tracking-system/domain"
	"github.com/jordanfduarte/vehicle-tracking-system/domain/repository"
	"github.com/jordanfduarte/vehicle-tracking-system/domain"
)

// VehicleRepositoryImpl Implements repository.VehicleRepository
type VehicleRepositoryImpl struct {
	Conn *gorm.DB
}

func VehicleRepositoryWithRDB(conn *gorm.DB) repository.VehiclesRepository {
	return &VehicleRepositoryImpl{Conn: conn}
}

func (r *VehicleRepositoryImpl) RemoveAll() error {
	return r.Conn.Exec( "DELETE FROM vehicles" ).Error
}

func (r *VehicleRepositoryImpl) GetAll() ([]domain.Vehicles, error) {
	// vehicles := []domain.Vehicles{}
	// if err := r.Conn.Preload("Vehicle").Find(&vehicles).Error; err != nil {
	// 	return nil, err
	// }

	// return vehicles, nil
	var vehicles []domain.Vehicles
	r.Conn.Raw(`
		SELECT
			v.Vehicle_ID,
			v.Fleet_ID,
			v.Name,
			IF (
				v.Max_Speed IS NULL,
				(SELECT fleets.Max_Speed FROM fleets WHERE fleets.Fleet_ID = v.Fleet_ID),
				v.Max_Speed
			) as Max_Speed
		FROM vehicles v
	`).Find(&vehicles)

	return vehicles, nil
}

func (r *VehicleRepositoryImpl) GetAllFleetAlertsByVehicle(id int) ([]domain.FleetAlerts, error) {
	var fleetAlerts []domain.FleetAlerts
	r.Conn.Raw(`
		SELECT
			fa.*
		FROM vehicles v
		INNER JOIN fleet_alerts fa ON v.Fleet_ID = fa.Fleet_ID
		WHERE v.Vehicle_ID = ?`, id).Find(&fleetAlerts)

	return fleetAlerts, nil
}

func (r *VehicleRepositoryImpl) Get(id int) (*domain.Vehicles, error) {
	vehicle := &domain.Vehicles{}
	// if err := r.Conn.Find(vehicle, "Vehicle_ID = ?", id).Error; err != nil {
	// 	return nil, err
	// }

	// return vehicle, nil

	r.Conn.Raw(`
		SELECT
			v.Vehicle_ID,
			v.Fleet_ID,
			v.Name,
			IF (
				v.Max_Speed IS NULL,
				(SELECT fleets.Max_Speed FROM fleets WHERE fleets.Fleet_ID = v.Fleet_ID),
				v.Max_Speed
			) as Max_Speed
		FROM vehicles v WHERE Vehicle_ID = ?`, id).Scan(&vehicle)

	return vehicle, nil
}

func (r *VehicleRepositoryImpl) Save(vehicle *domain.Vehicles) error {
	if err := r.Conn.Create(&vehicle).Error; err != nil {
		return err
	}

	return nil
}

func (r *VehicleRepositoryImpl) Remove(id int) error {
	return nil
}

func (r *VehicleRepositoryImpl) Update(*domain.Vehicles) error {
	return nil
}