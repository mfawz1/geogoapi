package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mfawz1/geogoapi/database"
	"github.com/stretchr/testify/assert"
)

func MakeRequest(req *http.Request) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	app := SetupAPI()
	w := httptest.NewRecorder()
	app.Router.ServeHTTP(w, req)
	return w
}
func TestFetchAllEntities(t *testing.T) {
	req, _ := http.NewRequest("GET", "/entities", nil)
	w := MakeRequest(req)
	assert.Equal(t, http.StatusOK, w.Code)
	var entities []database.GeoEntity
	app := SetupAPI()
	app.Db.Find(&entities)
	data, _ := json.Marshal(entities)
	t.Logf("data: %d", len(data))
	assert.Equal(t, string(data), w.Body.String())
}

func TestFetchSingleEntity(t *testing.T) {
	req, _ := http.NewRequest("GET", "/entities/408", nil)
	w := MakeRequest(req)
	assert.Equal(t, http.StatusOK, w.Code)
	var entity database.GeoEntity = database.GeoEntity{
		ID: 408,
	}
	app := SetupAPI()
	result := app.Db.Find(&entity)
	if result.RowsAffected == 0{
		data, _ := json.Marshal(map[string]string{})
		assert.Equal(t, string(data), w.Body.String())
		return
	}
	data, _ := json.Marshal(entity)
	assert.Equal(t, string(data), w.Body.String())
}
func TestFetchSingleEntityNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/entities/405", nil)
	w := MakeRequest(req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	var entity database.GeoEntity = database.GeoEntity{
		ID: 405,
	}
	app := SetupAPI()
	result := app.Db.Find(&entity)
	if result.RowsAffected == 0{
		data, _ := json.Marshal(map[string]string{})
		assert.Equal(t, string(data), w.Body.String())
		return
	}
	data, _ := json.Marshal(entity)
	assert.Equal(t, string(data), w.Body.String())
}
