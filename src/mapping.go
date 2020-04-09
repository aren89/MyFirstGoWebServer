package main

import (
	"MyFirstGoWebServer/src/jobApplication"
	"MyFirstGoWebServer/src/person"
	"github.com/gin-gonic/gin"
)

func CreateUrlMappings(router *gin.Engine) {
	v1 := router.Group("/api")

	person.Controller(v1.Group("/persons"))
	jobApplication.Controller(v1.Group("/job-applications"))
}
