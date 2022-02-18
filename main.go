package main

/**
* Initiating the server instance call
* @package main
* @author Jordan Duarte
* @versão 1.0.0
* @title da API Vehicle Tracking System
**/

import (
    "github.com/jordanfduarte/vehicle-tracking-system/interfaces"
)

func main() {
    // Change the port here
    interfaces.Run(8000)
}