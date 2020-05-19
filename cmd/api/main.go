package main

import (
	_applicationHandler "MyFirstGoWebServer/internal/application/handler"
	_applicationRepo "MyFirstGoWebServer/internal/application/repository"
	_applicationService "MyFirstGoWebServer/internal/application/service"
	"MyFirstGoWebServer/internal/database"
	_jobOfferHandler "MyFirstGoWebServer/internal/joboffer/handler"
	_jobOfferRepo "MyFirstGoWebServer/internal/joboffer/repository"
	_jobOfferService "MyFirstGoWebServer/internal/joboffer/service"
	_personHandler "MyFirstGoWebServer/internal/person/handler"
	_personRepo "MyFirstGoWebServer/internal/person/repository"
	_personService "MyFirstGoWebServer/internal/person/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("Startup api main")
	db, client := database.Init()
	router := gin.Default()
	api := router.Group("/api")

	personRepo := _personRepo.NewPersonRepository(db)
	personService := _personService.NewPersonService(personRepo)
	_personHandler.NewPersonHandler(api.Group("/persons"), personService)

	jobOfferRepo := _jobOfferRepo.NewJobOfferRepository(db)
	jobOfferService := _jobOfferService.NewJobOfferService(jobOfferRepo)
	_jobOfferHandler.NewJobOfferHandler(api.Group("/job-offers"), jobOfferService)

	applicationRepo := _applicationRepo.NewApplicationRepository(jobOfferRepo, db, client)
	applicationService := _applicationService.NewApplicationService(applicationRepo, personService)
	_applicationHandler.NewApplicationHandler(api.Group("/applications"), applicationService)

	router.Run(":8080")
}
