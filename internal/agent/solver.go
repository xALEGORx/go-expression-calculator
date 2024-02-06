package agent

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func Solver(messages <-chan amqp.Delivery) {
	for message := range messages {
		logrus.Infof("Received task: %s", message.Body)
	}
}
