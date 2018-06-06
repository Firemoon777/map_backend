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
