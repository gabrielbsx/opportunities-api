package router

import (
	"github.com/gabrielbsx/opportunities-api/docs"
	"github.com/gabrielbsx/opportunities-api/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		openingRoute := v1.Group("/openings")

		openingRoute.POST("", handler.CreateOpeningHandler)
		openingRoute.GET("", handler.ListOpeningsHandler)
		openingRoute.GET("/:id", handler.GetOpeningHandler)
		openingRoute.PUT("/:id", handler.UpdateOpeningHandler)
		openingRoute.DELETE("/:id", handler.DeleteOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
