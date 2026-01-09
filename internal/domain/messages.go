package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       uuid.UUID
	ThreadId uuid.UUID
	Content  string

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (m *Message) IsDeleted() bool {
	return m.DeletedAt != nil
}
