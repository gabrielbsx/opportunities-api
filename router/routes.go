package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Opening endpoint",
			})
		})
	}
}
