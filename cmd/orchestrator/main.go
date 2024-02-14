package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/agents"
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
	config_, err := config.Init()
	if err != nil {
		logrus.Fatalf("failed parse config .env: %s", err.Error())
		return
	}

	// try to connect to database
	if err = database.Init(); err != nil {
		logrus.Fatalf("database connection failed: %s", err.Error())
		return
	}
	if err = orchestrator.PrepareDatabase(); err != nil {
		logrus.Fatalf("database prepare sql failed: %s", err.Error())
		return
	}

	// try to connect to rabbitmq
	broker, err := rabbitmq.Init(config_.RabbitDSN)
	if err != nil {
		logrus.Fatal("rabbitmq connection failed")
		return
	}
	// try to create a queue for send task to rabbitmq
	if err = broker.InitQueue(config_.RabbitTaskQueue); err != nil {
		logrus.Fatal("rabbitmq fail creation a queue for tasks")
		return
	}
	// try to create a queue for server responses
	if err = broker.InitQueue(config_.RabbitAgentQueue); err != nil {
		logrus.Fatal("rabbitmq fail creation a queue for servers")
		return
	}

	// start listen a responses from agents
	messages, err := broker.ConnQueue(config_.RabbitAgentQueue)
	go agents.HandleAgentResponse(messages)
	go agents.HandleTimeoutAgents()
	go orchestrator.ResolveTasks()

	// initialization a gin
	gin.SetMode(config_.Mode)
	router := gin.Default()
	routes.InitRouter(router)

	logrus.Info("Orchestrator was successful started!")
	// run a server
	if err = router.Run(config_.ServerAddr); err != nil {
		logrus.Fatalf("failed run http server: %s", err.Error())
	}
}
