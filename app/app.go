package app

import "gorm.io/gorm"
import "github.com/gin-gonic/gin"
import databaseTypes "github.com/mfawz1/geogoapi/database"

func AppInit(db *gorm.DB) {
	// migrate models db.AutoMigrate()
	db.AutoMigrate(&databaseTypes.Entity{})
	router := gin.Default()
	//the router should only handle the RD part of CRUD, in relation to the geometry fields
	router.GET("/entity", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "hello data",
		})
	})
	router.GET("/getEntities", func(ctx *gin.Context){
		// fetch all entities 
		// add pagination
		// add query parameters
		ctx.JSON(200, gin.H{

		})
	})
	router.Run()
}

