package api

import (
	"github.com/gin-gonic/gin"
	databaseTypes "github.com/mfawz1/geogoapi/database"
)

func (app *AppAPI) entities() {
	app.Router.GET("/entities", func(ctx *gin.Context) {
		// fetch all entities
		// add pagination
		// add query parameters
		var entities []databaseTypes.GeoEntity
		app.Db.Find(&entities)
		ctx.JSON(200, entities)

	})
}
