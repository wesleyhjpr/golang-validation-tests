package main

import (
	"golang-gin-api-rest/controllers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestCheckStatusCodeOfTheGreeting(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controllers.Greeting)
	req, _ := http.NewRequest("GET", "/wesley", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "should be equal")
	mockResponse := `{"API says:":"what's up wesley"}`
	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, mockResponse, string(responseBody))
}
