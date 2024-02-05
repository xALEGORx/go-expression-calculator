package database

import (
	"github.com/jackc/pgx"

	"github.com/xALEGORx/go-expression-calculator/pkg/config"
)

var DB *pgx.Conn

func Init() error {
	var err error
	iconfig := config.Get()

	DB, err = pgx.Connect(pgx.ConnConfig{
		Host:     iconfig.PostgresHost,
		Port:     5432,
		Database: iconfig.PostgresDatabase,
		User:     iconfig.PostgresUser,
		Password: iconfig.PostgresPassword,
	})

	return err
}
