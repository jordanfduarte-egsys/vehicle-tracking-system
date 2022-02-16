package domain

import (
	"github.com/jinzhu/gorm"
)

// News represent entity of the Vehicle
type Vehicle struct {
    gorm.Model
    Vehicle_ID   int32    `json:"id" orm:"auto"`
    Fleet_ID     int32    `json:"fleet_id"`
    Name         string   `json:"name"`
    Max_Speed    string   `json:"max_speed"`
}