package main

import (
	"os"
	"fmt"
	"encoding/json"
)

func (conf *Configuration) Load(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Configuration error: ", err)
		return
	}
	
}
