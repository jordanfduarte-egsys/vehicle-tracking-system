package domain

/**
* Entity struct that represent mapping to data model
* @package domain
* @author Jordan Duarte
**/

// import (
//     "strconv"
//     "fmt"
//     "log"
// )

type VehiclePositions struct {
    Vehicle_Position_ID   int       `json:"id" gorm:"column:Vehicle_Position_ID;auto_increment;primary_key;not null"`
    Vehicle_ID            int       `json:"vehicle_id" gorm:"column:Vehicle_ID;type:int;not null"`
    Timestamp             string    `json:"timestamp" gorm:"column:Timestamp;type:varchar(255);not null"`
    Latitude              int       `json:"latitude" gorm:"column:Latitude;type:int;not null"`
    Longitude             int       `json:"longitude" gorm:"column:Longitude;type:int;not null"`
    Current_Speed         float32   `json:"current_speed" gorm:"column:Current_Speed;type:float(9,2);not null"`
    Max_Speed             float32   `json:"max_speed" gorm:"column:Max_Speed;type:float(9,2);not null"`
}

func NewVehiclePositions() FactoryDomain {
    return &VehiclePositions{}
}

func (a *VehiclePositions) IsValid() (isValid bool, error string) {
    isValid = true

    if a.Current_Speed < 0 {
        isValid = false
    }

    // _, err := strconv.ParseFloat(fmt.Sprintf("%v", a.Latitude), 32)
    // if err != nil {
    //     isValid = false
    // }

    // _, err2 := strconv.ParseFloat(fmt.Sprintf("%v", a.Longitude), 32)
    // if err2 != nil {
    //     isValid = false
    // }

    if (a.Timestamp != "ISO-8601") {
        a.Timestamp = "ISO-8601";
    }

    return isValid, ""
}