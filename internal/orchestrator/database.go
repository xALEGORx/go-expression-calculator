package orchestrator

import (
	"context"
	"github.com/xALEGORx/go-expression-calculator/pkg/database"
)

func PrepareDatabase() error {
	var sql []string = []string{
		"create table if not exists tasks (task_id serial primary key, expression text not null, status varchar(10) not null, answer text not null);",
		"create table if not exists agents (agent_id serial primary key, last_ping timestamp default CURRENT_TIMESTAMP);",
	}

	for _, query := range sql {
		if _, err := database.DB.Query(context.Background(), query); err != nil {
			return err
		}
	}

	return nil
}
