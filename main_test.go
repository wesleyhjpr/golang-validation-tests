package main

import (
	"bytes"
	"encoding/json"
	"golang-gin-api-rest/controllers"
	"golang-gin-api-rest/database"
	"golang-gin-api-rest/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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

	r.GET("/:name", controllers.Greeting)
	req, _ := http.NewRequest("GET", "/wesley", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "should be equal")
	responseMock := `{"API says:":"what's up wesley"}`
	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, responseMock, string(responseBody))
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

func TestGetStudentByCPFHandler(t *testing.T) {
	database.ConnectToDatabase()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)
	req, _ := http.NewRequest("GET", "/students/cpf/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudentByIDHandler(t *testing.T) {
	database.ConnectToDatabase()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/students/:id", controllers.GetStudentById)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var studentMock models.Student

	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, "Student test", studentMock.Name, "The name should be equal")
	assert.Equal(t, "12345678901", studentMock.CPF)
	assert.Equal(t, "123456789", studentMock.RG)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteStudentHandler(t *testing.T) {
	database.ConnectToDatabase()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUpdateStudentHandler(t *testing.T) {
	database.ConnectToDatabase()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.PATCH("/students/:id", controllers.UpdateStudent)

	student := models.Student{Name: "Student test",
		CPF: "99999999999", RG: "888888888"}
	jsonContent, _ := json.Marshal(student)

	path := "/students/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(jsonContent))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var studentUpdatedMock models.Student

	json.Unmarshal(response.Body.Bytes(), &studentUpdatedMock)

	assert.Equal(t, "99999999999", studentUpdatedMock.CPF)
	assert.Equal(t, "888888888", studentUpdatedMock.RG)
	assert.Equal(t, "Student test", studentUpdatedMock.Name)
}
