package schema

import "github.com/google/uuid"

type MessageCreate struct {
	ThreadId uuid.UUID `json:"thread_id" binding:"required"`
	Content  string    `json:"content" binding:"required"`
}

type MessagesByThreadRead struct {
	ThreadId uuid.UUID `json:"thread_id"`
}

type MessageRead struct {
	ID       uuid.UUID `json:"id"`
	ThreadId uuid.UUID `json:"thread_id"`
	Content  string    `json:"content"`
}
