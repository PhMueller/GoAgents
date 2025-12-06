package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

const (
	Driver       = "pgx"
	ErrorMessage = "Cannot connect to postgres db"
)

func Connect(ctx context.Context, databaseUrl string) *sqlx.DB {
	db, err := sqlx.ConnectContext(ctx, Driver, databaseUrl)

	if err != nil {
		panic(ErrorMessage)
	}

	return db
}
