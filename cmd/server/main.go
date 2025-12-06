package main

import (
	"context"
	"fmt"
	"net/http"

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

	messages := messageService.GetMessagesByThreadId(uuid.Must(uuid.Parse("69359037-9599-48e7-b8f2-48393c019135")))
	fmt.Printf("Messages: %+v\n", messages)

	s := server.NewServer()
	s.SetupRoutes(messageService)

	if err := http.ListenAndServe("localhost:8080", s.Router); err != nil {
		panic(err)
	}
}
