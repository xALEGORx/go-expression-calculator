package agents

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/services"
	"strconv"
)

func HandleProcessed(message amqp.Delivery) {
	// handle processed status from agent
	taskId, err := strconv.Atoi(message.CorrelationId)
	if err != nil {
		logrus.Errorf("Get wrong task id for proccessing: %s", message.CorrelationId)
		return
	}

	// set processed status into database
	if err = services.TaskService().SetProcessed(taskId, string(message.Body)); err != nil {
		logrus.Errorf("Failed update a row with task %d: %s", taskId, err.Error())
		return
	}

	logrus.Infof("Set status processed for #%s with agent #%s", message.CorrelationId, message.Body)
}
