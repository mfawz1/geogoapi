package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfawz1/geogoapi/database"
)

func (app *AppAPI) entities() {
	app.Router.GET("/entities", func(ctx *gin.Context) {
		// fetch all entities
		// add pagination
		// add query parameters
		var entities []database.GeoEntity
		app.Db.Find(&entities)
		ctx.JSON(http.StatusOK, entities)

	})
	app.Router.GET("/entities/:id", func(ctx *gin.Context){
		var entity database.GeoEntity = database.GeoEntity{}
		ctx.BindUri(&entity)
		result := app.Db.Find(&entity)
		if result.Error != nil{
			log.Panic(result.Error)
			ctx.JSON(http.StatusBadRequest, entity)
			return
		}
		if result.RowsAffected == 0{
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}
		ctx.JSON(http.StatusOK, entity)
	})
}
