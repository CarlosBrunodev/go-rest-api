package main

import (
	"github.com/CarlosBrunodev/go-rest-api/config"
	"github.com/CarlosBrunodev/go-rest-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.Getlogger("main")
	// Initializer configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	router.Initializer()
}
