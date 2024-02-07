package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/xALEGORx/go-expression-calculator/pkg/config"
)

var DB *pgxpool.Pool

func Init() error {
	var err error
	iconfig := config.Get()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", iconfig.PostgresUser, iconfig.PostgresPassword, iconfig.PostgresHost, iconfig.PostgresPort, iconfig.PostgresDatabase)

	DB, err = pgxpool.New(context.Background(), dsn)

	return err
}
