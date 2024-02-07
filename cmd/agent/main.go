package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/xALEGORx/go-expression-calculator/internal/agent"
	"github.com/xALEGORx/go-expression-calculator/pkg/logger"
	"github.com/xALEGORx/go-expression-calculator/pkg/rabbitmq"
)

type IConfig struct {
	WorkerID          string
	RabbitURL         string
	RabbitTaskQueue   string
	RabbitServerQueue string
	Threads           int
}

func main() {
	config := &IConfig{}
	flag.StringVar(&config.RabbitURL, "url", "amqp://guest:guest@localhost:5672", "RabbitMQ url for connection")
	flag.StringVar(&config.RabbitTaskQueue, "queue", "CalculatorTaskQueue1", "RabbitMQ queue name for listen")
	flag.StringVar(&config.RabbitServerQueue, "server", "CalculatorServerQueue1", "RabbitMQ queue name for server")
	flag.StringVar(&config.WorkerID, "worker", "worker", "Optional name of agent")
	flag.IntVar(&config.Threads, "threads", 5, "Threads count for goroutine")
	flag.Parse()

	// initialization a logrus
	logger.Init()

	// try to connect to rabbitmq
	broker, err := rabbitmq.Init(config.RabbitURL)
	if err != nil {
		logrus.Fatal("rabbitmq connection failed")
		return
	}
	messages, err := broker.ConnQueue(config.RabbitTaskQueue)
	done := make(chan bool)

	for i := 0; i < config.Threads; i++ {
		go agent.Solver(config.RabbitServerQueue, messages)
	}

	logrus.Infof("Agent \"%s\" was started with %d threads", config.WorkerID, config.Threads)
	<-done
}
