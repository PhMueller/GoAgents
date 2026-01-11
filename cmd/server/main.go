package main

import (
	"context"

	"examples.com/assistants/db/repository"
	"examples.com/assistants/internal/config"
	"examples.com/assistants/internal/db"
	"examples.com/assistants/internal/server"
	"examples.com/assistants/internal/services"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	conn := db.Connect(ctx, cfg.DatabaseURL)
	defer conn.Close(ctx)

	if conn.IsClosed() {
		panic("Failed to connect to the database")
	}

	queries := repository.New(conn)
	messageService := services.NewMessageService(queries)
	threadService := services.NewThreadService(queries)
	authService := services.NewAuthService(queries)

	s := server.NewServer()
	s.SetupRoutes(messageService, threadService, authService)
	s.AddValidators()

	if err := s.Engine.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
