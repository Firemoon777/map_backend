package main

import (
	//"fmt"
	"log"
	"net/http"
	//"encoding/json"
	//"strconv"


)

func main() {
	conf := Configuration{}
	conf.Load("conf.json")

	initDatabase(conf)
	router := initRouter(conf)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://207.154.254.134")
}

func StationHandler(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["stationId"], 10, 64)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	} 

	station := Station{Id: id, Name: []string{"Обводной канал", "Каретная"}, X: float32(5.0*id), Y: float32(5.0*id), Active: []bool{true, false}, Color: []int{0xFF0000, 0x654321}}
	json.NewEncoder(w).Encode(station)*/
}
