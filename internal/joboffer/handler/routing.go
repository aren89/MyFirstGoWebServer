package handler

import (
	"MyFirstGoWebServer/internal/core"
	"github.com/gin-gonic/gin"
)

type jobOfferHandlerImpl struct {
	JobOfferService core.JobOfferService
}

func NewJobOfferHandler(router *gin.RouterGroup, js core.JobOfferService) {
	handler := &jobOfferHandlerImpl{
		JobOfferService: js,
	}
	router.GET("/:id", handler.GetJobOfferDetail)
	router.GET("/", handler.GetJobOffers)
}