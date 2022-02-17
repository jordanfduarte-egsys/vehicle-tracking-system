package domain

import (
	"net/url"
)

// News represent entity of the FleetAlert
type FleetAlerts struct {
	Fleet_Alert_ID   int    `json:"id" gorm:"column:Fleet_Alert_ID;auto_increment;primary_key;not null"`
	Fleet_ID         int   	`json:"fleet_id" gorm:"column:Fleet_ID;type:int;not null"`
	WebHook          string `json:"webhook" gorm:"column:WebHook;type:varchar(255);not null"`
}

func (a *FleetAlerts) IsValid() (isValid bool) {
	isValid = true
	if a.WebHook == "" {
		isValid = false
	}
	_, err := url.ParseRequestURI(a.WebHook)
   	if err != nil {
		isValid = false
   	}

	u, err := url.Parse(a.WebHook)
   	if err != nil || u.Scheme == "" || u.Host == "" {
      	isValid = false
   	}
   	//fmt.Println(u)

	return isValid
}