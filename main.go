package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mfawz1/geogoapi/api"
	"github.com/mfawz1/geogoapi/database"
)

func main() {
	//init app
	_, testMode := os.LookupEnv("test_mode")
	if testMode {
		log.SetOutput(io.Discard)
	}
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	app := api.AppAPI{
		Router: router,
		Db:     database.SetupAndGetDB(),
	}
	app.API()
	app.Router.Run()
}
