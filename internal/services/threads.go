package services

import (
	"context"

	"examples.com/assistants/db/repository"
	"examples.com/assistants/internal/schema"
	"github.com/google/uuid"
)

type ThreadService struct {
	ctx     context.Context
	queries repository.Queries
}

func NewThreadService(ctx context.Context, queries repository.Queries) *ThreadService {
	threadService := ThreadService{
		ctx:     ctx,
		queries: queries,
	}
	return &threadService
}

func (t *ThreadService) CreateThread(threadCreate schema.CreateThreadRequest) (repository.Thread, error) {

	var title *string
	if threadCreate.Title != "" {
		title = &threadCreate.Title
	}

	dbThread, err := t.queries.CreateThread(t.ctx, title)
	return dbThread, err
}

func (t *ThreadService) GetThreadById(threadId uuid.UUID) (repository.Thread, error) {
	dbThread, err := t.queries.GetThreadById(t.ctx, threadId)
	return dbThread, err
}
