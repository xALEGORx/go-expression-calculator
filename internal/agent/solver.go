package agent

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/xALEGORx/go-expression-calculator/pkg/rabbitmq"
	"time"
)

func Solver(queueOrchestrator, agentId string, messages <-chan amqp.Delivery) {
	for message := range messages {
		if message.Type != "task" {
			continue
		}

		logrus.Infof("Received task #%s: %s", message.CorrelationId, message.Body)

		processed := amqp.Publishing{
			ContentType:   "text/plain",
			Body:          []byte(agentId),
			Type:          "processed",
			CorrelationId: message.CorrelationId,
		}
		if err := rabbitmq.Get().SendToQueue(queueOrchestrator, processed); err != nil {
			logrus.Fatalf("Failed sent status processed for task #%s: %s", message.CorrelationId, err.Error())
			continue
		}

		time.Sleep(5 * time.Second) // wait for response

		expression, err := govaluate.NewEvaluableExpression(string(message.Body))
		if err != nil {
			logrus.Errorf("Failed load expression task #%s: %s", message.CorrelationId, message.Body)
			continue
		}
		result, err := expression.Evaluate(nil)
		if err != nil {
			logrus.Errorf("Failed solve the expression %s for task #%s: %s", message.Body, message.CorrelationId, err.Error())
			continue
		}
		resultByte := []byte(fmt.Sprint(result))

		answer := amqp.Publishing{
			ContentType:   "text/plain",
			Body:          resultByte,
			Type:          "answer",
			CorrelationId: message.CorrelationId,
		}

		if err = rabbitmq.Get().SendToQueue(queueOrchestrator, answer); err != nil {
			logrus.Fatalf("Failed sent answer %s for task #%s: %s", resultByte, message.CorrelationId, err.Error())
			continue
		}

		logrus.Infof("Answer sent for task #%s: %s", message.CorrelationId, resultByte)
	}
}
