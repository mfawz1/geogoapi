package app

import (
	"github.com/gin-gonic/gin"
	databaseTypes "github.com/mfawz1/geogoapi/database"
	"gorm.io/gorm"
)

func AppInit(db *gorm.DB) {
	// migrate models db.AutoMigrate()
	db.AutoMigrate(&databaseTypes.GeoEntity{})
	router := gin.Default()
	//the router should only handle the RD part of CRUD, in relation to the geometry fields
	router.GET("/entity", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "hello data",
		})
	})
	// people who do APIs say nouns to represent resource is the sane default
	router.GET("/entities", func(ctx *gin.Context){
		// fetch all entities 
		// add pagination
		// add query parameters
		var entities []databaseTypes.Entity;
		db.Find(&entities);
		ctx.JSON(200, entities)
	})
	router.Run()
}

