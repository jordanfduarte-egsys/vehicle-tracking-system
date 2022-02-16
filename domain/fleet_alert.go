package domain

import (
	"github.com/jinzhu/gorm"
)

// News represent entity of the Fleet_Alert
type Fleet_Alert struct {
	gorm.Model
	Fleet_Alert_ID   int    `json:"id" orm:"auto"`
	Fleet_ID         int   	`json:"fleet_id"`
	WebHook          string `josn:"web_hook"`
}