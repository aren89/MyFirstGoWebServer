package person

import (
	"MyFirstGoWebServer/src/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Controller(router *gin.RouterGroup) {
	router.GET("/:id", GetPersonDetail)
	router.GET("/", GetPerson)
	router.POST("/", PostPerson)
	router.PUT("/:id", UpdatePerson)
}

func GetPerson(c *gin.Context) {
	c.JSON(404, gin.H{"error": "person not found"})
}

func GetPersonDetail(c *gin.Context) {
	c.JSON(404, gin.H{"error": "person not found"})
}

func UpdatePerson(c *gin.Context) {
	c.JSON(404, gin.H{"error": "person not found"})
}

func PostPerson(c *gin.Context) {
	var person PersonModel
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personId, err := ManagePersonToSave(person)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	serializer := common.PostSerializer{
		C:  c,
		Id: personId,
	}
	c.JSON(http.StatusCreated, serializer.Response())

}