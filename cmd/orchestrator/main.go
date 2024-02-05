package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/routes"
	"github.com/xALEGORx/go-expression-calculator/pkg/config"
	"github.com/xALEGORx/go-expression-calculator/pkg/database"
	"github.com/xALEGORx/go-expression-calculator/pkg/logger"
)

// @title Expression Calculator
// @version 1.0
// @description Endpoints for expression calculator by ALEGOR
// @BasePath /api/v1
func main() {
	logger.Init()
	config_ := config.Init()

	if err := database.Init(); err != nil {
		logrus.Fatal("database connection failed")
		return
	}

	gin.SetMode(config_.Mode)

	router := gin.Default()
	routes.InitRouter(router)

	err := router.Run(config_.ServerAddr)
	if err != nil {
		logrus.Fatal("database connection failed")
		return
	}
}
