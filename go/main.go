package main

import (
	"github.com/IsaqueAmorim/DevHunter/config"
	"github.com/IsaqueAmorim/DevHunter/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Initialize()
}
