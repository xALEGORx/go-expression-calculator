package services

import "github.com/xALEGORx/go-expression-calculator/internal/orchestrator/repositories"

type Task struct {
	repo *repositories.Task
}

func (t *Task) Create(expression string) (repositories.TaskModel, error) {
	taskId, err := t.repo.Create(expression)

	if err != nil {
		return repositories.TaskModel{}, err
	}

	task := repositories.TaskModel{
		TaskID:     taskId,
		Expression: expression,
	}

	return task, nil
}

// create new task service
func TaskService() *Task {
	return &Task{}
}
