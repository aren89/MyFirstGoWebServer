package main

import (
	"MyFirstGoWebServer/internal/database"
	_personHandler "MyFirstGoWebServer/internal/person/handler"
	_personRepo "MyFirstGoWebServer/internal/person/repository"
	_personService "MyFirstGoWebServer/internal/person/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Init()
	router := gin.Default()
	api := router.Group("/api")

	personRepo := _personRepo.NewPersonRepository(db)
	personService := _personService.NewPersonService(personRepo)
	_personHandler.NewPersonHandler(api.Group("/persons"), personService)

	router.Run(":8080")
}
