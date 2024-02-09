package agents

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func HandleAgentResponse(messages <-chan amqp.Delivery) {
	if err := InitAgents(); err != nil {
		logrus.Errorf("Failed loading a agent list: %s", err.Error())
		return
	}

	for message := range messages {
		if message.Type == "answer" {
			HandleAnswer(message)
		}
		if message.Type == "processed" {
			HandleProcessed(message)
		}
		if message.Type == "ping" {
			HandlePing(message)
		}
	}
}
