package main

import (
	"golang-gin-api-rest/controllers"
	"golang-gin-api-rest/database"
	"golang-gin-api-rest/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "Student test",
		CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestCheckStatusCodeOfTheGreeting(t *testing.T) {
	r := SetupTestRoutes()

	CreateStudentMock()
	defer DeleteStudentMock()

	r.GET("/:name", controllers.Greeting)
	req, _ := http.NewRequest("GET", "/wesley", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "should be equal")
	mockResponse := `{"API says:":"what's up wesley"}`
	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, mockResponse, string(responseBody))
}

func TestGetAllStudentsHandler(t *testing.T) {
	database.ConnectToDatabase()
	r := SetupTestRoutes()
	r.GET("/students", controllers.GetAllStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudentByCPF(t *testing.T) {
	database.ConnectToDatabase()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/students/cpf/:cpf", controllers.GetAllStudents)
	req, _ := http.NewRequest("GET", "/students/cpf/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}
