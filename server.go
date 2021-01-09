package main

import (
	"net/http"

	"github.com/Wesleyss071299/Golang-rest-api/controller"
	"github.com/Wesleyss071299/Golang-rest-api/middlewares"
	"github.com/Wesleyss071299/Golang-rest-api/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := VideoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":8080")
}
