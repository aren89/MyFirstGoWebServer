package handler

import (
	"MyFirstGoWebServer/internal/core"
	"github.com/gin-gonic/gin"
)

type personHandlerImpl struct {
	PersonService core.PersonService
}

func NewPersonHandler(router *gin.RouterGroup, ps core.PersonService) {
	handler := &personHandlerImpl{
		PersonService: ps,
	}
	router.POST("/", handler.PostPerson)
	router.GET("/:id", handler.GetPersonDetail)
	router.GET("/", handler.GetPersons)
}
