package repositories

import "github.com/xALEGORx/go-expression-calculator/pkg/database"

type Task struct {
}

type TaskModel struct {
	TaskID     int    `json:"task_id"`
	Expression string `json:"expression"`
}

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
		if err = rows.Scan(&task.TaskID, &task.Expression); err != nil {
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

	if err := database.DB.QueryRow("INSERT INTO tasks (expression) VALUES ($1) RETURNING task_id", expression).Scan(&insertedID); err != nil {
		return 0, err
	}

	return insertedID, nil
}

func TaskRepository() *Task {
	return &Task{}
}
