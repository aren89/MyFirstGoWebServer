package handler

import (
	"MyFirstGoWebServer/internal/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *applicationHandlerImpl) DeleteApplication(c *gin.Context) {
	applicationId := c.Param("id")
	err := h.ApplicationService.Delete(c, applicationId)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *applicationHandlerImpl) PostApplication(c *gin.Context) {
	var application core.ApplicationRepresentation
	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	applicationId, err := h.ApplicationService.SaveApplication(c, &application)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	response := core.Post{Id: applicationId}
	c.JSON(http.StatusCreated, response)
}

func (h *applicationHandlerImpl) GetApplicationDetail(c *gin.Context) {
	applicationId := c.Param("id")
	applicationRepresentation, err := h.ApplicationService.GetById(c, applicationId)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	if "" == applicationRepresentation.Id {
		c.JSON(http.StatusNotFound, gin.H{"error": "NO_APPLICATION_FOUND"})
		return
	}
	c.JSON(http.StatusCreated, applicationRepresentation)
}
