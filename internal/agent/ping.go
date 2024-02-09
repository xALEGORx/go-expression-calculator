package agent

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/xALEGORx/go-expression-calculator/pkg/rabbitmq"
	"time"
)

func Ping(queueOrchestrator string, agentId string, pingTime int) {
	ticket := time.NewTicker(time.Duration(pingTime) * time.Second)

	for {
		select {
		case <-ticket.C:
			answer := amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(agentId),
				Type:        "ping",
			}

			if err := rabbitmq.Get().SendToQueue(queueOrchestrator, answer); err != nil {
				logrus.Fatalf("Failed sent ping: %s", err.Error())
				break
			}

			logrus.Debugf("Ping was successful sent")
		}
	}
}
