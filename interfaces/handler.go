package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"regexp"
	"strconv"

	"../application"
	"../config"
	"../domain"
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

	// Executa a migração inicial dos dados
	migrate()

	return r
}

// =============================
//    ACTIONS
// =============================

func deleteAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// err = application.RemoveFleetAlertAll()
	// if err != nil {
	// 	Error(w, http.StatusNotFound, err, err.Error())
	// 	return
	// }

	// err = application.RemoveFleetAll()
	// if err != nil {
	// 	Error(w, http.StatusNotFound, err, err.Error())
	// 	return
	// }

	// err = application.VehiclePositionAll()
	// if err != nil {
	// 	Error(w, http.StatusNotFound, err, err.Error())
	// 	return
	// }

	err = application.RemoveVehicleAll()
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, nil)
}

// index action
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	JSON(w, http.StatusOK, "Api is running!")
}

// func getNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	param := ps.ByName("param")

// 	// if param is numeric than search by news_id, otherwise
// 	// if alphabetic then search by topic.Slug
// 	newsID, err := strconv.Atoi(param)
// 	if err != nil {
// 		// param is alphabetic
// 		news, err2 := application.GetNewsByTopic(param)
// 		if err2 != nil {
// 			Error(w, http.StatusNotFound, err2, err2.Error())
// 			return
// 		}

// 		JSON(w, http.StatusOK, news)
// 		return
// 	}

// 	// param is numeric
// 	news, err := application.GetNews(newsID)
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	JSON(w, http.StatusOK, news)
// }

// func getAllNews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	queryValues := r.URL.Query()
// 	status := queryValues.Get("status")

// 	// if status parameter exist draft|deleted|publish
// 	if status == "draft" || status == "deleted" || status == "publish" {
// 		news, err := application.GetAllNewsByFilter(status)
// 		if err != nil {
// 			Error(w, http.StatusNotFound, err, err.Error())
// 			return
// 		}

// 		JSON(w, http.StatusOK, news)
// 		return
// 	}

// 	limit := queryValues.Get("limit")
// 	page := queryValues.Get("page")

// 	// if custom pagination exist news?limit=15&page=2
// 	if limit != "" && page != "" {
// 		limit, _ := strconv.Atoi(limit)
// 		page, _ := strconv.Atoi(page)

// 		if limit != 0 && page != 0 {
// 			news, err := application.GetAllNews(limit, page)
// 			if err != nil {
// 				Error(w, http.StatusNotFound, err, err.Error())
// 				return
// 			}

// 			JSON(w, http.StatusOK, news)
// 			return
// 		}
// 	}

// 	news, err := application.GetAllNews(15, 1) // 15, 1 default pagination
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	JSON(w, http.StatusOK, news)
// }

// func createNews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	decoder := json.NewDecoder(r.Body)
// 	var p domain.News
// 	if err := decoder.Decode(&p); err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 	}

// 	err := application.AddNews(p)
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	JSON(w, http.StatusCreated, nil)
// }

// func removeNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	newsID, err := strconv.Atoi(ps.ByName("news_id"))
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	err = application.RemoveNews(newsID)
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	JSON(w, http.StatusOK, nil)
// }

// func updateNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	decoder := json.NewDecoder(r.Body)
// 	var p domain.News
// 	err := decoder.Decode(&p)
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 	}

// 	newsID, err := strconv.Atoi(ps.ByName("news_id"))
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	err = application.UpdateNews(p, newsID)
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	JSON(w, http.StatusOK, nil)
// }

// =============================
//    TOPIC
// =============================

// func getTopic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	topicID, err := strconv.Atoi(ps.ByName("topic_id"))
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	topic, err := application.GetTopic(topicID)
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	JSON(w, http.StatusOK, topic)
// }

// func getAllTopic(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	topics, err := application.GetAllTopic()
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	JSON(w, http.StatusOK, topics)
// }

// func createTopic(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

// 	type payload struct {
// 		Name string `json:"name"`
// 		Slug string `json:"slug"`
// 	}
// 	var p payload
// 	err := json.NewDecoder(r.Body).Decode(&p)
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	err = application.AddTopic(p.Name, p.Slug)
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	JSON(w, http.StatusCreated, nil)
// }

// func removeTopic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	topicID, err := strconv.Atoi(ps.ByName("topic_id"))
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	err = application.RemoveTopic(topicID)
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	JSON(w, http.StatusOK, nil)
// }

// func updateTopic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	decoder := json.NewDecoder(r.Body)
// 	var p domain.Topic
// 	err := decoder.Decode(&p)
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 	}

// 	topicID, err := strconv.Atoi(ps.ByName("topic_id"))
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	err = application.UpdateTopic(p, topicID)
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err, err.Error())
// 		return
// 	}

// 	JSON(w, http.StatusOK, nil)
// }

// =============================
//    MIGRATE
// =============================

func migrate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := config.DBMigrate()
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}
}