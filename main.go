package main

import (
	"io"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/mfawz1/geogoapi/api"
	"github.com/mfawz1/geogoapi/database"
	"github.com/mfawz1/geogoapi/log"
)

func main() {
	log.InitLoggers()
	//init app
	_, testMode := os.LookupEnv("test_mode")
	if testMode {
		log.InfoLog.SetOutput(io.Discard)
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
