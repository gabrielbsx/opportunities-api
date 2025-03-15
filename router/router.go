package router

import (
	"github.com/gabrielbsx/opportunities-api/config"
	"github.com/gin-gonic/gin"
)

var (
	logger *config.Logger
)

func Initialize() {
	// Initialize router
	logger = config.GetLogger("router")

	router := gin.Default()

	err := router.SetTrustedProxies(nil)

	if err != nil {
		logger.Error("Error setting trusted proxies: ", err)
		return
	}

	// Initialize routes
	initializeRoutes(router)

	// Run the server
	router.Run(":3001")
}
