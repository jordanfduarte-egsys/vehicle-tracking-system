package domain

/**
* Entity struct that represent mapping to data model
* @package domain
* @author Jordan Duarte
**/

type Fleets struct {
    Fleet_ID   int      `json:"id" gorm:"column:Fleet_ID;auto_increment;primary_key;not null"`
    Name       string   `json:"name" gorm:"column:Name;type:varchar(255);not null"`
    Max_Speed  float32  `json:"max_speed" gorm:"column:Max_Speed;type:float(9,2);not null"`
}

func NewFleets() FactoryDomain {
    return &Fleets{}
}

func (a *Fleets) IsValid() (isValid bool, error string) {
    isValid = true
    if a.Name == "" {
        error = "A velocidade deve ser maior que 0."
        isValid = false
    }

    if len(a.Name) == 0 || len(a.Name) > 254 {
        error = "A velocidade deve ser maior que 0."
        isValid = false
    }

    if a.Max_Speed <= 0 {
        error = "A velocidade deve ser maior que 0."
        isValid = false
    }
    return isValid, error
}