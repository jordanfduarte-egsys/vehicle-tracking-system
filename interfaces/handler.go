package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"strconv"
	"github.com/jordanfduarte/vehicle-tracking-system/application"
	"github.com/jordanfduarte/vehicle-tracking-system/config"
	"github.com/jordanfduarte/vehicle-tracking-system/domain"
)

// Run start server
func Run(port int) error {
	log.Printf("Server running at http://localhost:%d/", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
	r := httprouter.New()

	/**
	* Init Route
	*/
	r.GET("/", indexAction)

	/**
	* Reset
	*/
	// limpa toda a base de dados
	r.DELETE("/database", deleteAction)

	/**
	* Fleets
	*/
	// lista todas as frotas
	r.GET("/fleets", fleetsGetAction)
	// Cria uma frota
	r.POST("/fleets", fleetsPostAction)
	// lista todas os alertas de uma frota
	r.GET("/fleets/:id/alerts", alertsGetAction)
	// Cria uma alerta para frota
	r.POST("/fleets/:id/alerts", alertsPostAction)

	/**
	* Vehicles
	*/
	// lista todos os veículos
	r.GET("/vehicles", vehiclesGetAction)
	// Cria uma veículo
	r.POST("/vehicles", vehiclesPostAction)
	/*
	* Vehicles Positions
	*/
	// lista todos as posições de um veículo
	r.GET("/vehicles/:id/positions", positionsGetAction)
	// Salva a posição do veículo)
	r.POST("/vehicles/:id/positions", positionsPostAction)

	// Migrations
	r.GET("/migration", migrationAction)

	return r
}

// =============================
//    ACTIONS
// =============================

func deleteAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := application.RemoveVehiclePositionAll()
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err2 := application.RemoveVehicleAll()
	if err2 != nil {
		Error(w, http.StatusNotFound, err2, err2.Error())
		return
	}

	err3 := application.RemoveFleetAlertAll()
	if err3 != nil {
		Error(w, http.StatusNotFound, err3, err3.Error())
		return
	}

	err4 := application.RemoveFleetAll()
	if err4 != nil {
		Error(w, http.StatusNotFound, err4, err4.Error())
		return
	}

	JSON(w, http.StatusOK, "Delete has been processed")
}

func fleetsGetAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fleets, err := application.GetAllFleets()

	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, fleets)
}
// Cria uma frota
func fleetsPostAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var fleet *domain.Fleets
	err := json.NewDecoder(r.Body).Decode(&fleet)

    if err != nil {
        Error(w, http.StatusBadRequest, err, "Invalid parameters entered")
        return
    }

	// str := fmt.Sprintf("%v", fleet.Max_Speed)
	// log.Printf(str)
	isValid, error := fleet.IsValid()
	// log.Printf(isValid)
	if isValid == false {
		Error(w, http.StatusBadRequest, nil, error)
		return
	}

	//criar
	err2 := application.AddFleet(fleet)
	if err2 != nil {
		Error(w, http.StatusNotFound, err2, err2.Error())
		return
	}

	type ReturnDinamyc struct {
		Id int `json:"id"`
	}
	returnDinamyc := &ReturnDinamyc{Id: fleet.Fleet_ID}
	JSON(w, http.StatusCreated, returnDinamyc)
}
// lista todas os alertas de uma frota
func alertsGetAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := ps.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	// veririca se exist o Fleet
	fleet, err := application.GetRowFleet(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	if fleet.Fleet_ID <= 0 {
		Error(w, http.StatusNotFound, nil, "")
		return
	}

	fleetAlerts, err := application.GetAllFleetAlerts(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, fleetAlerts)
}
// Cria uma alerta para frota
func alertsPostAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := ps.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	var fleetAlert *domain.FleetAlerts
	err2 := json.NewDecoder(r.Body).Decode(&fleetAlert)

    if err2 != nil {
        Error(w, http.StatusBadRequest, err2, "Invalid parameters entered")
        return
    }

	// veririca se exist o Fleet
	fleet, err := application.GetRowFleet(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	if fleet.Fleet_ID <= 0 {
		Error(w, http.StatusNotFound, nil, "")
		return
	}

	// str := fmt.Sprintf("%v", fleet.Max_Speed)
	// log.Printf(str)
	isValid := fleetAlert.IsValid()
	// log.Printf(isValid)
	if isValid == false {
		Error(w, http.StatusBadRequest, nil, "")
		return
	}

	//criar
	fleetAlert.Fleet_ID = id
	err3 := application.AddFleetAlert(fleetAlert)
	if err3 != nil {
		Error(w, http.StatusNotFound, err3, err3.Error())
		return
	}

	// verificar aqui ......
	//{ "webhook": "http://localhost:8081/fleet/alert" }
	//defaultStruct := &domain.DefaultStruct{Id: fleet.Fleet_ID}
	type ReturnDinamyc struct {
		WebHook string `json:"webhook"`
	}
	returnTypeDinamyc := &ReturnDinamyc{WebHook: fleetAlert.WebHook}

	JSON(w, http.StatusCreated, returnTypeDinamyc)
}

func vehiclesGetAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// veririca se exist o Fleet
	vehicles, err := application.GetAllVehicles()
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, vehicles)
}

func vehiclesPostAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var vehicle *domain.Vehicles
	err2 := json.NewDecoder(r.Body).Decode(&vehicle)

    if err2 != nil {
        Error(w, http.StatusBadRequest, err2, "Invalid parameters entered")
        return
    }

	// str := fmt.Sprintf("%v", fleet.Max_Speed)
	// log.Printf(str)
	isValid := vehicle.IsValid()
	// log.Printf(isValid)
	if isValid == false {
		Error(w, http.StatusBadRequest, nil, "")
		return
	}

	fleet, err := application.GetRowFleet(vehicle.Fleet_ID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	if fleet.Fleet_ID <= 0 {
		Error(w, http.StatusNotFound, nil, "")
		return
	}

	//criar
	err3 := application.AddVehicle(vehicle)
	if err3 != nil {
		Error(w, http.StatusNotFound, err3, err3.Error())
		return
	}

	type ReturnDinamyc struct {
		Id int `json:"id"`
	}
	returnTypeDinamyc := &ReturnDinamyc{Id: vehicle.Vehicle_ID}

	JSON(w, http.StatusCreated, returnTypeDinamyc)
}

func positionsGetAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := ps.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	// veririca se exist o Fleet
	positions, err := application.GetAllPositionsByVehicles(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, positions)
}

func positionsPostAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := ps.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	var vehiclePosition *domain.VehiclePositions
	err2 := json.NewDecoder(r.Body).Decode(&vehiclePosition)

    if err2 != nil {
        Error(w, http.StatusBadRequest, err2, "Invalid parameters entered")
        return
    }

	// veririca se exist o Fleet
	vehicle, err := application.GetRowVehicle(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	if vehicle.Vehicle_ID <= 0 {
		Error(w, http.StatusNotFound, nil, "")
		return
	}

	// str := fmt.Sprintf("%v", fleet.Max_Speed)
	// log.Printf(str)
	isValid := vehiclePosition.IsValid()
	// log.Printf(isValid)
	if isValid == false {
		log.Printf("dfdfd")
		Error(w, http.StatusBadRequest, nil, "")
		return
	}

	//criar
	vehiclePosition.Vehicle_ID = id
	vehicleFind, err65 := application.GetRowVehicle(id)

	if err65 != nil {
		Error(w, http.StatusNotFound, err65, err65.Error())
		return
	}

	vehiclePosition.Max_Speed = vehicleFind.Max_Speed
	err3 := application.AddPositionVehicle(vehiclePosition)
	if err3 != nil {
		Error(w, http.StatusNotFound, err3, err3.Error())
		return
	}

	// depois de salvo
	if vehiclePosition.Current_Speed > vehiclePosition.Max_Speed {
		// busca todos os webhooks cadastrados na frota
		fleetAlerts, err5 := application.GetAllFleetAlertsByVehicle(id)
		if err5 != nil {
			Error(w, http.StatusNotFound, err5, err5.Error())
			return
		}

		//Worker Pools
		jobs := make(chan domain.Site, 100)
		for w := 1; w <= 3; w++ {
			go domain.Crawl(w, jobs)
		}

		for _, fleetAlertSend := range fleetAlerts {
			jsonValue, _ := json.Marshal(vehiclePosition)
			//log.Println(fleetAlertSend.WebHook)
			jobs <- domain.Site{URL: fleetAlertSend.WebHook, Buffer: jsonValue}
		}
		close(jobs)
	}

	type ReturnDinamyc struct {
		Id int `json:"id"`
	}
	returnTypeDinamyc := &ReturnDinamyc{Id: vehiclePosition.Vehicle_Position_ID}

	JSON(w, http.StatusCreated, returnTypeDinamyc)
}

// index action
func indexAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	JSON(w, http.StatusOK, "Api is running!")
}

// =============================
//    MIGRATE
// =============================

func migrationAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := config.DBMigrate()
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	// application.AddFleet(&domain.Fleet{Name: "Veículos de perseguição", Max_Speed: 30.55});
	// application.AddFleet(&domain.Fleet{Name: "Veículos de transporte de prisioneiros", Max_Speed: 25});
	// application.AddFleet(&domain.Fleet{Name: "Escolta armada", Max_Speed: 22.22});

	log.Println("Migration has been processed")
	JSON(w, http.StatusOK, "Migration has been processed!")
}