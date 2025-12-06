package domain

import (
	"github.com/google/uuid"
)

type Message struct {
	ID       uuid.UUID
	ThreadId uuid.UUID
	Content  string
}
