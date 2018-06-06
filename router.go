package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func initRouter(conf Configuration) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/events/", EventHandler)
	//router.HandleFunc("/station/{stationId}", StationHandler)
	router.HandleFunc("/map/colors", ColorHandler)
	router.HandleFunc("/map/events", EventHandler)
	return router
}

func setupOrigin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://207.154.254.134")
}

func ColorHandler(w http.ResponseWriter, r *http.Request) {
	setupOrigin(w, r)
	var c []Color
	db.Find(&c)
	json.NewEncoder(w).Encode(c);
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	setupOrigin(w, r)
	var e []Event
	db.Order("event_time ASC", true)
	db.Find(&e)
	json.NewEncoder(w).Encode(e);
}
