package agents

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/repositories"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/services"
	"strconv"
)

func HandleError(message amqp.Delivery) {
	// handle error from agent
	taskId, err := strconv.Atoi(message.CorrelationId)
	if err != nil {
		logrus.Errorf("Get wrong task id for error: %s", message.CorrelationId)
		return
	}

	// set error into database
	if err = services.TaskService().SetAnswer(taskId, string(message.Body), repositories.STATUS_FAIL); err != nil {
		logrus.Errorf("Failed update a row with task %d: %s", taskId, err.Error())
		return
	}

	logrus.Infof("Get error for %s: %s", message.CorrelationId, message.Body)
}
