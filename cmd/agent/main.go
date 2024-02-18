package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/xALEGORx/go-expression-calculator/internal/agent"
	"github.com/xALEGORx/go-expression-calculator/pkg/logger"
	"github.com/xALEGORx/go-expression-calculator/pkg/rabbitmq"
)

type IConfig struct {
	AgentID          string
	RabbitURL        string
	RabbitTaskQueue  string
	RabbitAgentQueue string
	Threads          int
	Ping             int
	Wait             int
	Debug            bool
}

func main() {
	config := &IConfig{}
	flag.StringVar(&config.RabbitURL, "url", "amqp://guest:guest@localhost:5672", "RabbitMQ url for connection")
	flag.StringVar(&config.RabbitTaskQueue, "queue", "CalculatorTaskQueue1", "RabbitMQ queue name for listen")
	flag.StringVar(&config.RabbitAgentQueue, "server", "CalculatorAgentQueue1", "RabbitMQ queue name for agents")
	flag.StringVar(&config.AgentID, "agent", "agent", "Name ID of agent")
	flag.IntVar(&config.Threads, "threads", 5, "Threads count for goroutine")
	flag.IntVar(&config.Ping, "ping", 60, "Ping time in seconds")
	flag.IntVar(&config.Wait, "wait", 5, "Wait time (emulating long query)")
	flag.BoolVar(&config.Debug, "debug", false, "Enable debug mode")
	flag.Parse()

	// initialization a logrus
	logger.Init(config.Debug)

	// try to connect to rabbitmq
	broker, err := rabbitmq.Init(config.RabbitURL)
	if err != nil {
		logrus.Fatalf("rabbitmq connection failed: %s", err.Error())
		return
	}
	messages, err := broker.ConnQueue(config.RabbitTaskQueue)
	done := make(chan bool)

	// start threads for solving tasks
	for i := 0; i < config.Threads; i++ {
		go agent.Solver(config.RabbitAgentQueue, config.AgentID, config.Wait, messages)
	}
	// start ping method
	go agent.Ping(config.RabbitAgentQueue, config.AgentID, config.Ping)

	logrus.Infof("Agent \"%s\" was started with %d threads", config.AgentID, config.Threads)
	<-done
}
