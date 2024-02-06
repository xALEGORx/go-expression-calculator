package repositories

import "github.com/xALEGORx/go-expression-calculator/pkg/database"

type Task struct {
}

type TaskModel struct {
	TaskID     int    `json:"task_id"`
	Expression string `json:"expression"`
	Status     string `json:"status"`
	Error      string `json:"error"`
}

const (
	STATUS_CREATED   = "created"
	STATUS_PROGRESS  = "progress"
	STATUS_FAIL      = "fail"
	STATUS_COMPLETED = "completed"
)

// Get all tasks in database
func (t *Task) GetAllTasks() ([]TaskModel, error) {
	rows, err := database.DB.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []TaskModel{}

	for rows.Next() {
		var task TaskModel
		if err = rows.Scan(&task.TaskID, &task.Expression, &task.Status, &task.Error); err != nil {
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

	query := "INSERT INTO tasks (expression, status, error) VALUES ($1, $2, $3) RETURNING task_id"
	if err := database.DB.QueryRow(query, expression, STATUS_CREATED, "").Scan(&insertedID); err != nil {
		return 0, err
	}

	return insertedID, nil
}

// Find task by id
func (t *Task) GetById(taskId int) (TaskModel, error) {
	var task TaskModel

	query := "SELECT * FROM tasks WHERE task_id = $1"
	if err := database.DB.QueryRow(query, taskId).Scan(&task.TaskID, &task.Expression, &task.Status, &task.Error); err != nil {
		return task, err
	}

	return task, nil
}

func TaskRepository() *Task {
	return &Task{}
}
