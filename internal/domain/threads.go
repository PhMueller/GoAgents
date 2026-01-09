package domain

import "time"

type Thread struct {
	ID    string
	Title *string

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (t *Thread) IsDeleted() bool {
	return t.DeletedAt != nil
}
