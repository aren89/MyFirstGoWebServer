package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	CreateUrlMappings(router)

	router.Run(":8080")
}
