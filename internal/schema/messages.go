package schema

import "github.com/google/uuid"

type MessageCreate struct {
	ThreadId uuid.UUID `json:"thread_id" binding:"required"`
	Content  string    `json:"content" binding:"required"`
}

type MessageRead struct {
	ID       string `json:"id"`
	ThreadId string `json:"thread_id"`
	Content  string `json:"content"`
}
