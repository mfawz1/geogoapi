package main

import (
	"log"
	"github.com/mfawz1/geogoapi/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func main() {
	db, err := gorm.Open(postgres.Open(loadDatabaseConfig()), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect to db", err)
	}
    db.Exec("create extension if not exists postgis")
	log.Print("Database connected successfully!")
	//init app
	app.AppInit(db)
}
