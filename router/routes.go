package router

import (
	"github.com/gabrielbsx/opportunities-api/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		openingRoute := v1.Group("/openings")

		openingRoute.POST("", handler.CreateOpeningHandler)
		openingRoute.GET("", handler.ListOpeningsHandler)
		openingRoute.GET("/:id", handler.GetOpeningHandler)
		openingRoute.PUT("/:id", handler.UpdateOpeningHandler)
		openingRoute.DELETE("/:id", handler.DeleteOpeningHandler)
	}
}
