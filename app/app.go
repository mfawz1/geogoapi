package app

import (
	"github.com/gin-gonic/gin"
	api "github.com/mfawz1/geogoapi/api"
	databaseTypes "github.com/mfawz1/geogoapi/database"
	"gorm.io/gorm"
)

func AppInit(db *gorm.DB) {
	// migrate models db.AutoMigrate()
	db.AutoMigrate(&databaseTypes.GeoEntity{})
	router := gin.Default()
	app := api.AppAPI{
		Router: router,
		Db:     db,
	}
	app.API()
}
