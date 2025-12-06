package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

const (
	ErrorMessage = "Cannot connect to postgres db"
)

func Connect(ctx context.Context, databaseUrl string) *pgx.Conn {
	conn, err := pgx.Connect(ctx, databaseUrl)
	if err != nil {
		panic(ErrorMessage)
	}
	return conn
}
