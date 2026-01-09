package main

import (
	"context"
	"fmt"

	"examples.com/assistants/db/repository"
	"examples.com/assistants/internal/config"
	"examples.com/assistants/internal/db"
	"examples.com/assistants/internal/server"
	"examples.com/assistants/internal/services"
	"github.com/google/uuid"
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
	messageService := services.NewMessageService(ctx, *queries)
	threadService := services.NewThreadService(ctx, *queries)

	messages := messageService.GetMessagesByThreadID(uuid.Must(uuid.Parse("69359037-9599-48e7-b8f2-48393c019135")))
	fmt.Printf("Messages: %+v\n", messages)

	s := server.NewServer()
	s.SetupRoutes(messageService, threadService)
	s.AddValidators()

	if err := s.Engine.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
