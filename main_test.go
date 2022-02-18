package main

/**
* Endpoint test case. Using the go test -v command in the bash console
* @package main
* @author Jordan Duarte
* @versão 1.0.0
* @title da API Vehicle Tracking System
**/

import (
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "testing"
    "bytes"
    "github.com/julienschmidt/httprouter"
    "github.com/jordanfduarte/vehicle-tracking-system/domain"
    "github.com/jordanfduarte/vehicle-tracking-system/action"
)

// Test endpoint FleetsGet
func TestFleetsGetAction(t *testing.T) {
    req, err := http.NewRequest(http.MethodGet, "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        action.NewFleetHandler().FleetsGetAction(w, r, httprouter.Params{})
    })
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}

// Test endpoint FleetsPost
func TestFleetsPostAction(t *testing.T) {
    fleet := &domain.Fleets{
        Name: "Veículos de perseguição",
        Max_Speed: 30.55,
    }

    var jsonStr, _ = json.Marshal(fleet)
    req, err := http.NewRequest(http.MethodPost, "/fleets", bytes.NewBuffer(jsonStr))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "Content-Type: text/plain")
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        action.NewFleetHandler().FleetsPostAction(w, r, httprouter.Params{})
    })
    handler.ServeHTTP(rr, req)

    status := rr.Code;
    if status == http.StatusBadRequest {
        t.Error("Invalid parameters entered")
    } else if status != http.StatusCreated {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }

    var fleetReturned *domain.Fleets
    err2 := json.NewDecoder(rr.Body).Decode(&fleetReturned)
    if err2 != nil {
        t.Errorf("Handler returned unexpected error: %v", err2.Error())
    }

    if fleetReturned.Fleet_ID <= 0 {
        t.Errorf("Handler returned unexpected object id: %v", fleetReturned.Fleet_ID)
    }
}

// Test endpoint AlertsGet
func TestAlertsGetAction(t *testing.T) {
    req, err := http.NewRequest(http.MethodGet, "/fleets/{id}/alerts", nil)
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        action.NewAlertsHandler().AlertsGetAction(w, r, httprouter.Params{
            httprouter.Param{
                Key: "id",
                Value: "1",
            },
        })
    })
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}

// Test endpoint AlertsPost
func TestAlertsPostAction(t *testing.T) {
    fleetAlerts := &domain.FleetAlerts{
        WebHook: "http://localhost:8081/fleet/alert",
    }

    var jsonStr, _ = json.Marshal(fleetAlerts)
    req, err := http.NewRequest(http.MethodPost, "/fleets/{id}/alerts", bytes.NewBuffer(jsonStr))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        action.NewAlertsHandler().AlertsPostAction(w, r, httprouter.Params{
            httprouter.Param{
                Key: "id",
                Value: "1",
            },
        })
    })
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status == http.StatusBadRequest {
        t.Error("Invalid parameters entered")
    } else if status := rr.Code; status == http.StatusNotFound {
        t.Error("Fleet not found")
    } else if status := rr.Code; status != http.StatusCreated {
        t.Errorf("Handler returned wrong status code: got %v want %v",
            status, http.StatusCreated)
    }

    var fleetAlertsReturn *domain.FleetAlerts
    err2 := json.NewDecoder(rr.Body).Decode(&fleetAlertsReturn)
    if err2 != nil {
        t.Fatal(err2)
    }

    if fleetAlertsReturn.Fleet_Alert_ID <= 0 {
        t.Errorf("Handler returned unexpected object id: %v", fleetAlertsReturn.Fleet_Alert_ID)
    }
}

// Test endpoint VehiclesGet
func TestVehiclesGetAction(t *testing.T) {
    req, err := http.NewRequest(http.MethodGet, "/vehicles", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        action.NewVehicleHandler().VehiclesGetAction(w, r, httprouter.Params{})
    })
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}

// Test endpoint VehiclesPost
func TestVehiclesPostAction(t *testing.T) {
    vehicle := &domain.Vehicles{
        Fleet_ID: 1,
        Name: "veículo 1",
        Max_Speed: 50,
    }

    var jsonStr, _ = json.Marshal(vehicle)
    req, err := http.NewRequest(http.MethodPost, "/vehicles", bytes.NewBuffer(jsonStr))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        action.NewVehicleHandler().VehiclesPostAction(w, r, httprouter.Params{})
    })
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status == http.StatusBadRequest {
        t.Error("Invalid parameters entered")
    } else if status := rr.Code; status != http.StatusCreated {
        t.Errorf("Handler returned wrong status code: got %v want %v",
            status, http.StatusCreated)
    }

    var vehicleReturned *domain.Vehicles
    err2 := json.NewDecoder(rr.Body).Decode(&vehicleReturned)
    if err2 != nil {
        t.Fatal(err2)
    }

    if vehicleReturned.Vehicle_ID <= 0 {
        t.Errorf("Handler returned unexpected object id: %v", vehicleReturned.Vehicle_ID)
    }
}

// Test endpoint PositionsGet
func TestPositionsGetAction(t *testing.T) {
    req, err := http.NewRequest(http.MethodGet, "/vehicles/{id}/positions", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        action.NewVehiclePositionHandler().PositionsGetAction(w, r, httprouter.Params{
            httprouter.Param{
                Key: "id",
                Value: "1",
            },
        })
    })
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status == http.StatusNotFound {
        t.Error("Vehicles not found")
    } else if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}

// Test endpoint PositionsPost
func TestPositionsPostAction(t *testing.T) {
    vehiclePosition := &domain.VehiclePositions{
        Timestamp: "ISO-8601",
        Latitude: 0,
        Longitude: 0,
        Current_Speed: 0,
    }

    var jsonStr, _ = json.Marshal(vehiclePosition)
    req, err := http.NewRequest(http.MethodPost, "/vehicles/{id}/positions", bytes.NewBuffer(jsonStr))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        action.NewVehiclePositionHandler().PositionsPostAction(w, r, httprouter.Params{
            httprouter.Param{
                Key: "id",
                Value: "1",
            },
        })
    })
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status == http.StatusBadRequest {
        t.Error("Invalid parameters entered")
    } else if status := rr.Code; status == http.StatusNotFound {
        t.Error("Vehicles not found")
    } else if status := rr.Code; status != http.StatusCreated {
        t.Errorf("Handler returned wrong status code: got %v want %v",
            status, http.StatusCreated)
    }

    var vehiclePositionReturned *domain.VehiclePositions
    err2 := json.NewDecoder(rr.Body).Decode(&vehiclePositionReturned)
    if err2 != nil {
        t.Fatal(err2)
    }

    if vehiclePositionReturned.Vehicle_Position_ID <= 0 {
        t.Errorf("Handler returned unexpected object id: %v", vehiclePositionReturned.Vehicle_Position_ID)
    }
}

// Test endpoint Index
func TestIndexAction(t *testing.T) {
    req, err := http.NewRequest(http.MethodGet, "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        action.NewIndexHandler().IndexAction(w, r, httprouter.Params{})
    })
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `{"message":"Api is running!"}`
    if rr.Body.String() != expected {
        t.Errorf("Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }

    if rr.Code != http.StatusOK {
        t.Errorf("Handler returned unexpected code status http: got %v want %v",
            rr.Code, http.StatusOK)
    }
}

