package repositories

import (
	"context"
	"github.com/xALEGORx/go-expression-calculator/pkg/database"
)

type Task struct {
}

type TaskModel struct {
	TaskID     int    `json:"task_id"`
	Expression string `json:"expression"`
	Status     string `json:"status"`
	Answer     string `json:"answer"`
	AgentID    string `json:"agent_id"`
}

const (
	STATUS_CREATED   = "created"
	STATUS_PROCESSED = "processed"
	STATUS_FAIL      = "fail"
	STATUS_COMPLETED = "completed"
)

// Get all tasks in database
func (t *Task) GetAllTasks() ([]TaskModel, error) {
	rows, err := database.DB.Query(context.Background(), "SELECT * FROM tasks ORDER BY task_id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []TaskModel{}

	for rows.Next() {
		var task TaskModel
		if err = rows.Scan(&task.TaskID, &task.Expression, &task.Status, &task.Answer, &task.AgentID); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Create new row with task
func (t *Task) Create(expression string) (int, error) {
	var insertedID int

	query := "INSERT INTO tasks (expression, status, answer, agent_id) VALUES ($1, $2, $3, $4) RETURNING task_id"
	if err := database.DB.QueryRow(context.Background(), query, expression, STATUS_CREATED, "", "").Scan(&insertedID); err != nil {
		return 0, err
	}

	return insertedID, nil
}

// Find task by id
func (t *Task) GetById(taskId int) (TaskModel, error) {
	var task TaskModel

	query := "SELECT * FROM tasks WHERE task_id = $1"
	if err := database.DB.QueryRow(context.Background(), query, taskId).Scan(&task.TaskID, &task.Expression, &task.Status, &task.Answer, &task.AgentID); err != nil {
		return task, err
	}

	return task, nil
}

// Update answer for expression by task id
func (t *Task) SetAnswer(taskId int, answer string, status string) error {
	query := "UPDATE tasks SET answer = $1, status = $2 WHERE task_id = $3"
	if _, err := database.DB.Exec(context.Background(), query, answer, status, taskId); err != nil {
		return err
	}

	return nil
}

// Update status to processed and set agent_id
func (t *Task) SetProcessed(taskId int, agentId string) error {
	query := "UPDATE tasks SET status = $1, agent_id = $2 WHERE task_id = $3"
	if _, err := database.DB.Exec(context.Background(), query, STATUS_PROCESSED, agentId, taskId); err != nil {
		return err
	}

	return nil
}

func TaskRepository() *Task {
	return &Task{}
}
