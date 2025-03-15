package main

import (
	"github.com/gabrielbsx/opportunities-api/config"
	"github.com/gabrielbsx/opportunities-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("Opportunities API")

	// Initialize configs
	err := config.Init()

	if err != nil {
		logger.Error("Error initializing configs: ", err)
		return
	}

	// Initialize routers
	router.Initialize()
}
