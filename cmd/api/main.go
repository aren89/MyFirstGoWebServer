package main

import (
	"MyFirstGoWebServer/internal/database"
	_jobOfferRepo "MyFirstGoWebServer/internal/joboffer/repository"
	_jobOfferService "MyFirstGoWebServer/internal/joboffer/service"
	_jobOfferHandler "MyFirstGoWebServer/internal/joboffer/handler"
	_personHandler "MyFirstGoWebServer/internal/person/handler"
	_personRepo "MyFirstGoWebServer/internal/person/repository"
	_personService "MyFirstGoWebServer/internal/person/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("Startup api main")
	db := database.Init()
	router := gin.Default()
	api := router.Group("/api")

	personRepo := _personRepo.NewPersonRepository(db)
	personService := _personService.NewPersonService(personRepo)
	_personHandler.NewPersonHandler(api.Group("/persons"), personService)

	jobOfferRepo := _jobOfferRepo.NewJobOfferRepository(db)
	jobOfferService := _jobOfferService.NewJobOfferService(jobOfferRepo)
	_jobOfferHandler.NewJobOfferHandler(api.Group("/job-offers"), jobOfferService)

	router.Run(":8080")
}
