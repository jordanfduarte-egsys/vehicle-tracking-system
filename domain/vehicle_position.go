package domain

import (
    "github.com/jinzhu/gorm"
)

// News represent entity of the Vehicle_Position
type Vehicle_Position struct {
    gorm.Model
    Vehicle_Position_ID   int32    `json:"id" orm:"auto"`
    Vehicle_ID            int32    `json:"fleet_id"`
    Timestamp             int64    `json:"timestamp"`
    Latitude              int32    `json:"latitude"`
    Longitude             int32    `json:"longitude"`
    Current_Speed         float32  `json:"current_speed"`
    Max_Speed             float32  `json:"max_speed"`
}