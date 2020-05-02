package handler

import (
	"MyFirstGoWebServer/internal/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *personHandlerImpl) GetPersonDetail(c *gin.Context) {
	personId := c.Param("id")
	personRepresentation, err := h.PersonService.GetById(c, personId)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	if (core.PersonRepresentation{}) == personRepresentation {
		c.JSON(http.StatusNotFound, gin.H{"error": "NO_PERSON_FOUND"})
		return
	}
	c.JSON(http.StatusCreated, personRepresentation)
}

func (h *personHandlerImpl) GetPersons(c *gin.Context) {
	emailFilter := c.Query("email")
	yearsOfExperienceWorkingFilter := c.Query("yearsOfExperienceWorking")
	estimatedLevelFilter := c.Query("estimatedLevel")
	var result = make([]core.PersonRepresentation, 0)
	persons, err := h.PersonService.Fetch(c, emailFilter, yearsOfExperienceWorkingFilter, estimatedLevelFilter)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	result = append(result, persons...)
	c.JSON(http.StatusCreated, result)
}

func (h *personHandlerImpl) PostPerson(c *gin.Context) {
	var person core.PersonRepresentation
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personId, err := h.PersonService.StoreWithEstimatedLevel(c, &person)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	response := core.Post{Id: personId}
	c.JSON(http.StatusCreated, response)
}
