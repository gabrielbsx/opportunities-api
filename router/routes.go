package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	// Ping test
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
