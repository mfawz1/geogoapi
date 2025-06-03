package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppAPI struct {
	Router *gin.Engine
	Db     *gorm.DB
}

// register api end point
func (app *AppAPI) API() {
	// the router should only handle the RD part of CRUD, in relation to the geometry fields, perhaps ID as well?
	// people who do APIs say nouns to represent resource is the sane default
	app.entities()
	app.Router.Run()
}
