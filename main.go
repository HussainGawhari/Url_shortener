package main

import (
	"url_shortener/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/link/:code", controller.GetShortLink)
	router.GET("/:code", controller.GetShortLink)
	router.POST("/link", controller.GenerateShortLink)
	router.DELETE("/link/:code", controller.Delete)
	router.Run(":8080")

}
