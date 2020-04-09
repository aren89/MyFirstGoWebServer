package jobApplication

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Controller(router *gin.RouterGroup) {
	router.GET("/:id", GetJobApplicationDetail)
	router.POST("/", PostJobApplicationPerson)
	router.DELETE("/:id", CancelJobApplication)

}

func GetJobApplicationDetail(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Job Application not found"})

}

func PostJobApplicationPerson(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Job Application not found"})
}

func CancelJobApplication(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Job Application not found"})
}
