package orchestrator

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/repositories"
	"strconv"
)

func HandleServerResponse(messages <-chan amqp.Delivery) {
	for message := range messages {
		if message.Type == "answer" {
			// handle answer from agent
			taskId, err := strconv.Atoi(message.CorrelationId)
			if err != nil {
				logrus.Errorf("Get wrong task id for answer: %s", message.CorrelationId)
				continue
			}

			if err := repositories.TaskRepository().SetAnswer(taskId, string(message.Body), repositories.STATUS_COMPLETED); err != nil {
				logrus.Errorf("Failed update a row with task %d: %s", taskId, err.Error())
				continue
			}

			logrus.Printf("Get answer for %s: %s", message.CorrelationId, message.Body)
		}
		if message.Type == "ping" {
			// TODO: handle ping
		}
	}
}
