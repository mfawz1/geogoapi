package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mfawz1/geogoapi/database"
	"gorm.io/gorm"
)

type AppAPI struct {
	Router *gin.Engine
	Db     *gorm.DB
}

func SetupAPI() AppAPI {
	//init app
	router := gin.Default()
	app := AppAPI{
		Router: router,
		Db:     database.SetupAndGetDB(),
	}
	app.API()
	return app
}

// register api end point
func (app *AppAPI) API() {
	app.Db.AutoMigrate(&database.GeoEntity{})
	// the router should only handle the RD part of CRUD, in relation to the geometry fields, perhaps ID as well?
	// people who do APIs say nouns to represent resource is the sane default
	app.entities()
}
