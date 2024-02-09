package orchestrator

import (
	"context"
	"github.com/xALEGORx/go-expression-calculator/pkg/database"
)

func PrepareDatabase() error {

	var sql = []string{
		"set time zone 'Europe/Moscow'",
		"create table if not exists tasks (task_id serial primary key, expression text not null, status varchar(10) not null, answer text not null);",
		"create table if not exists agents (agent_id varchar(255) primary key, last_ping timestamp with time zone default CURRENT_TIMESTAMP);",
	}

	for _, query := range sql {
		if _, err := database.DB.Query(context.Background(), query); err != nil {
			return err
		}
	}

	return nil
}
