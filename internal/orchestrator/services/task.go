package services

import (
	"github.com/streadway/amqp"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/repositories"
	"github.com/xALEGORx/go-expression-calculator/pkg/config"
	"github.com/xALEGORx/go-expression-calculator/pkg/rabbitmq"
	"github.com/xALEGORx/go-expression-calculator/pkg/websocket"
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

	// return task model for response

	task := repositories.TaskModel{
		TaskID:     taskId,
		Status:     repositories.STATUS_CREATED,
		Expression: expression,
	}

	// send message to websocket

	wsData := websocket.WSData{
		Action: "new_task",
		Id:     taskId,
		Data:   task,
	}
	if err = websocket.Broadcast(wsData); err != nil {
		return repositories.TaskModel{}, err
	}

	return task, nil
}

// create new task service
func TaskService() *Task {
	return &Task{}
}
