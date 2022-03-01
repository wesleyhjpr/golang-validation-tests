package main

import (
	"golang-gin-api-rest/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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

	if response.Code != http.StatusOK {
		t.Fatalf("Status error: the amount received was %d and what was expected was %d", response.Code, http.StatusOK)
	}
}
