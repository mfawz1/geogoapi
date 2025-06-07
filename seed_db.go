package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/mfawz1/geogoapi/database"
	"github.com/mfawz1/geogoapi/log"
)

//	func main() {
//		SeedDatabase()
//	}
func SeedDatabase() {
	db := database.SetupAndGetDB()
	db.AutoMigrate(&database.GeoEntity{})
	for range 100 {
		var entity database.GeoEntity
		gofakeit.Struct(&entity)
		entity.Lat = gofakeit.Latitude()
		entity.Lng = gofakeit.Longitude()
		log.InfoLog.Printf("setting Lat: %f", entity.Lat)
		log.InfoLog.Printf("setting Lng: %f", entity.Lng)
		entity.Name = gofakeit.Name()
		if err := db.Create(&entity).Error; err != nil {
			log.ErrorLog.Fatal(err)
		}
	}
}
