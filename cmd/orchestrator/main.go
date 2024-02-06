package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/routes"
	"github.com/xALEGORx/go-expression-calculator/pkg/config"
	"github.com/xALEGORx/go-expression-calculator/pkg/database"
	"github.com/xALEGORx/go-expression-calculator/pkg/logger"
	"github.com/xALEGORx/go-expression-calculator/pkg/rabbitmq"
)

// @title Expression Calculator
// @version 1.0
// @description Endpoints for expression calculator by ALEGOR
// @BasePath /api/v1
func main() {
	// initialization a logrus
	logger.Init()
	// parsing .env file to config struct
	config_ := config.Init()

	// try to connect to database
	if err := database.Init(); err != nil {
		logrus.Fatal("database connection failed")
		return
	}

	// try to connect to rabbitmq
	broker, err := rabbitmq.Init(config_.RabbitURL, config_.RabbitQueue)
	if err != nil {
		logrus.Fatal("rabbitmq connection failed")
		return
	}
	// try to create a queue for rabbitmq
	if err = broker.InitQueue(); err != nil {
		logrus.Fatal("rabbitmq fail creation a queue")
		return
	}

	// initialization a gin
	gin.SetMode(config_.Mode)
	router := gin.Default()
	routes.InitRouter(router)

	logrus.Info("Orchestrator was successful started!")
	// run a server
	if err = router.Run(config_.ServerAddr); err != nil {
		logrus.Fatal("database connection failed")
		return
	}
}
