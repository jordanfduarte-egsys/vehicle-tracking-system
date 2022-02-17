package domain


// News represent entity of the VehiclePosition
type VehiclePositions struct {
    Vehicle_Position_ID   int       `json:"id" gorm:"column:Vehicle_Position_ID;auto_increment;primary_key;not null"`
    Vehicle_ID            int       `json:"vehicle_id" gorm:"column:Vehicle_ID;type:int;not null"`
    Timestamp             string    `json:"timestamp" gorm:"column:Timestamp;type:varchar(255);not null"`
    Latitude              float32   `json:"latitude" gorm:"column:Latitude;type:decimal(10, 8);not null"`
    Longitude             float32   `json:"longitude" gorm:"column:Longitude;type:decimal(10, 8);not null"`
    Current_Speed         float32   `json:"current_speed" gorm:"column:Current_Speed;type:float(9,2);not null"`
    Max_Speed             float32   `json:"max_speed" gorm:"column:Max_Speed;type:float(9,2);not null"`
}


func (a *VehiclePositions) IsValid() (isValid bool) {
	isValid = true

    if a.Current_Speed < 0 {
		isValid = false
	}

    if (a.Timestamp != "ISO-8601") {
        a.Timestamp = "ISO-8601";
    }

	return isValid
}