package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func initDatabase(conf Configuration) {
	t := conf.User + ":" + conf.Password + "@tcp(" + conf.Url  + ")/metro_map"
	db, err := gorm.Open("mysql", t)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Station{})
	db.AutoMigrate(&Color{})
	db.AutoMigrate(&Event{})
}
