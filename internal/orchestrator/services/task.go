package services

import (
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/repositories"
	"github.com/xALEGORx/go-expression-calculator/pkg/rabbitmq"
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
	if err = rabbitmq.Get().SendTask(expression); err != nil {
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
