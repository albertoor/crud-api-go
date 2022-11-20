package main

import (
	"example/crud-api-go/models"
	"example/crud-api-go/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Student{})

	router := routes.SetupStudentsRoutes(db)
	router.Run()
}
