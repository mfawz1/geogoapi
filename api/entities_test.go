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

func TestEntitesAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := SetupAPI()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/entities", nil)
	app.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	entities := make([]database.GeoEntity, 0)
	data, _ := json.Marshal(entities)
	t.Logf("data: %s", string(data))
	assert.Equal(t, string(data), w.Body.String())

}
