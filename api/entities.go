package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfawz1/geogoapi/database"
	"github.com/mfawz1/geogoapi/log"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetEntitiesInRange(Db *gorm.DB, point database.GeoPoint, distance float64) []database.GeoEntity {
	var entities []database.GeoEntity
	result := Db.Where("st_distance(st_transform(geo_point, 3857), st_transform(?::geometry, 3857)) < ?", point.String(), fmt.Sprintf("%f", distance)).Find(&entities)
	if result.Error != nil {
		log.ErrorLog.Print(result.Error)
	}
	return entities
}

func fetchEntitiesWithLatLngFilter(app *AppAPI) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// fetch all entities
		// add pagination
		// add query parameters
		lat := ctx.DefaultQuery("lat", "")
		lng := ctx.DefaultQuery("lng", "")
		distance := ctx.DefaultQuery("distance", "100")
		if len(lat) > 0 && len(lng) > 0 {
			lat64, lat64err := strconv.ParseFloat(lat, 64)
			lng64, lng64err := strconv.ParseFloat(lng, 64)
			dist64, disterr := strconv.ParseFloat(distance, 64)
			if lat64err == nil && lng64err == nil && disterr == nil {
				log.InfoLog.Print("Fetching lat & lng & distance: ", lat64, lng64, dist64)
				point := database.GeoPoint{
					Lat: lat64,
					Lng: lng64,
				}
				entities := GetEntitiesInRange(app.Db, point, dist64)
				ctx.JSON(http.StatusOK, entities)
				return
			} else {
				log.ErrorLog.Fatal(lat64err, lng64err, disterr)
			}
		}
		var entities []database.GeoEntity
		app.Db.Find(&entities)
		ctx.JSON(http.StatusOK, entities)
		return

	}
}
func fetcEntitiesWithId(app *AppAPI) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var entity database.GeoEntity = database.GeoEntity{}
		ctx.BindUri(&entity)
		result := app.Db.Find(&entity)
		if result.Error != nil {
			log.ErrorLog.Panic(result.Error)
			ctx.JSON(http.StatusBadRequest, entity)
			return
		}
		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}
		ctx.JSON(http.StatusOK, entity)
	}
}
func (app *AppAPI) entities() {
	app.Router.GET("/entities", fetchEntitiesWithLatLngFilter(app))
	app.Router.GET("/entities/:id", fetcEntitiesWithId(app))
}
