package domain

import (
	"github.com/jinzhu/gorm"
)

// News represent entity of the Fleet
type Fleet struct {
	gorm.Model
	Fleet_ID   int32    `json:"id" orm:"auto"`
	Name       string  `json:"name"`
	Max_Speed  float32  `json:"max_speed"`
}