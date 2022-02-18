package domain

/**
* Factory for create dynamic struct
* @package domain
* @author Jordan Duarte
**/

type FactoryDomain interface {
    IsValid() (bool, string)
}

func GetFactoryDomain(name string) FactoryDomain {
    if name == "fleet" {
        return NewFleets()
    }

    if name == "fleetAlert" {
        return NewFleetAlerts()
    }

    if name == "vehicle" {
        return NewVehicles()
    }

    if name == "vehiclePosition" {
        return NewVehiclePositions()
    }

    return nil
}