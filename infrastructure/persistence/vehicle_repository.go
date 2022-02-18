package persistence

/**
* Implements repository interface
* @package persistence
* @author Jordan Duarte
**/

import (
    "gorm.io/gorm"
    "github.com/jordanfduarte/vehicle-tracking-system/domain/repository"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
)

type VehicleRepositoryImpl struct {
    Conn *gorm.DB
}

func VehicleRepositoryWithRDB(conn *gorm.DB) repository.VehiclesRepository {
    return &VehicleRepositoryImpl{Conn: conn}
}

func (r *VehicleRepositoryImpl) RemoveAll() error {
    e := r.Conn.Exec("DELETE FROM vehicles").Error
    if e != nil {
        return e
    }
    return r.Conn.Exec("ALTER TABLE vehicles AUTO_INCREMENT=0;").Error
}

func (r *VehicleRepositoryImpl) GetAll() ([]domain.Vehicles, error) {
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

func (r *VehicleRepositoryImpl) GetFirst() (*domain.Vehicles, error) {
    vehicles := &domain.Vehicles{}
    r.Conn.Raw(`
        SELECT * FROM vehicles ORDER BY Vehicle_ID ASC LIMIT 1`).Find(&vehicles)

    return vehicles, nil
}

func (r *VehicleRepositoryImpl) Save(vehicleCheckNullParam *repository.VehicleCheckNullParam) error {
    if vehicleCheckNullParam.IsNullMaxSpeed {
        sqlStatement  := `INSERT INTO vehicles (Fleet_ID, Name, Max_Speed) VALUES (?, ?, NULL)`
        r.Conn.Exec(
            sqlStatement,
            vehicleCheckNullParam.Vehicles.Fleet_ID,
            vehicleCheckNullParam.Vehicles.Name)

        r.Conn.Raw(`
            SELECT * FROM vehicles ORDER BY Vehicle_ID DESC LIMIT 1`).Scan(&vehicleCheckNullParam.Vehicles)
    } else {
        if err := r.Conn.Create(&vehicleCheckNullParam.Vehicles).Error; err != nil {
            return err
        }
    }

    return nil
}
