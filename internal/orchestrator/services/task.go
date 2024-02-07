package services

import (
	"github.com/streadway/amqp"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/repositories"
	"github.com/xALEGORx/go-expression-calculator/pkg/config"
	"github.com/xALEGORx/go-expression-calculator/pkg/rabbitmq"
	"strconv"
)

type Task struct {
	repo *repositories.Task
}

func (t *Task) Create(expression string) (repositories.TaskModel, error) {
	taskId, err := t.repo.Create(expression)

	if err != nil {
		return repositories.TaskModel{}, err
	}

	// send task to queue of rabbitmq

	message := amqp.Publishing{
		ContentType:   "text/plain",
		Body:          []byte(expression),
		Type:          "task",
		CorrelationId: strconv.Itoa(taskId),
	}

	if err = rabbitmq.Get().SendToQueue(config.Get().RabbitTaskQueue, message); err != nil {
		return repositories.TaskModel{}, err
	}

	task := repositories.TaskModel{
		TaskID:     taskId,
		Status:     repositories.STATUS_CREATED,
		Expression: expression,
	}

	return task, nil
}

// create new task service
func TaskService() *Task {
	return &Task{}
}
