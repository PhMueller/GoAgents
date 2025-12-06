package main

import (
	"context"
	"net/http"

	"examples.com/assistants/internal/config"
	"examples.com/assistants/internal/db"
	"examples.com/assistants/internal/repositories"
	"examples.com/assistants/internal/server"
	"examples.com/assistants/internal/services"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	postgresDatabase := db.Connect(ctx, cfg.DatabaseURL)
	messageRepository := repositories.NewMessageRepository(postgresDatabase)
	messageService := services.NewMessageService(messageRepository)

	s := server.NewServer()
	s.SetupRoutes(messageService)

	if err := http.ListenAndServe("localhost:8080", s.Router); err != nil {
		panic(err)
	}
}
