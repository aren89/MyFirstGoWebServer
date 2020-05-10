package handler

import (
	"MyFirstGoWebServer/internal/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *jobOfferHandlerImpl) GetJobOfferDetail(c *gin.Context) {
	jobOfferId := c.Param("id")
	jobOfferRepresentation, err := h.JobOfferService.GetById(c, jobOfferId)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	if (core.JobOfferRepresentation{}) == jobOfferRepresentation {
		c.JSON(http.StatusNotFound, gin.H{"error": "NO_JOB_OFFER_FOUND"})
		return
	}
	c.JSON(http.StatusCreated, jobOfferRepresentation)
}

func (h *jobOfferHandlerImpl) GetJobOffers(c *gin.Context) {
	roleFilter := c.Query("role")
	companyFilter := c.Query("company")
	var result = make([]core.JobOfferRepresentation, 0)
	jobOffers, err := h.JobOfferService.Fetch(c, roleFilter, companyFilter)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	result = append(result, jobOffers...)
	c.JSON(http.StatusCreated, result)
}