package domain

// News represent entity of the Vehicle
type Vehicles struct {
    Vehicle_ID   int      `json:"id" gorm:"column:Vehicle_ID;auto_increment;primary_key;not null"`
    Fleet_ID     int      `json:"fleet_id" gorm:"column:Fleet_ID;type:int;not null"`
    Name         string   `json:"name" gorm:"column:Name;type:varchar(255);not null"`
    Max_Speed    float32  `json:"max_speed" gorm:"column:Max_Speed;type:float(9,2)"`
}


func (a *Vehicles) IsValid() (isValid bool) {
	isValid = true
	if a.Fleet_ID <= 0 {
		isValid = false
	}

	if len(a.Name) == 0 || len(a.Name) > 254 {
		isValid = false
	}

	if a.Max_Speed < 0 {
		isValid = false
	}

	return isValid
}