package main

import (
	"MyFirstGoWebServer/src/common"
	"github.com/gin-gonic/gin"
)

func main() {
	common.Init()
	router := gin.Default()
	CreateUrlMappings(router)
	router.Run(":8080")
}
