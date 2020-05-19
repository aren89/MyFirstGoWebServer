package handler

import (
	"MyFirstGoWebServer/internal/core"
	"github.com/gin-gonic/gin"
)

type applicationHandlerImpl struct {
	ApplicationService core.ApplicationService
}

func NewApplicationHandler(router *gin.RouterGroup, as core.ApplicationService) {
	handler := &applicationHandlerImpl{
		ApplicationService: as,
	}
	router.POST("/", handler.PostApplication)
	router.DELETE("/:id", handler.DeleteApplication)
	router.GET("/:id", handler.GetApplicationDetail)
}
