package main

import (
	"log"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/mfawz1/geogoapi/database"
)

// func main() {
// 	SeedDatabase()
// }
func SeedDatabase() {
	db := database.SetupAndGetDB()
	db.AutoMigrate(&database.GeoEntity{})
	for range 100 {
		var entity database.GeoEntity
		gofakeit.Struct(&entity)
		entity.Lat = gofakeit.Latitude()
		entity.Lng = gofakeit.Longitude()
		log.Printf("setting Lat: %f", entity.Lat)
		log.Printf("setting Lng: %f", entity.Lng)
		entity.Name = gofakeit.Name()
		if err := db.Create(&entity).Error; err != nil {
			log.Fatal(err)
		}
	}
}
