package main

import (
	"golang-gin-api-rest/database"
	"golang-gin-api-rest/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}
