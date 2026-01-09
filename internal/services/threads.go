package services

import (
	"context"

	"examples.com/assistants/db/repository"
	"examples.com/assistants/internal/domain"
	"examples.com/assistants/internal/schema"
	"github.com/google/uuid"
)

type ThreadService struct {
	ctx     context.Context
	queries repository.Queries
}

func NewThreadService(ctx context.Context, queries repository.Queries) *ThreadService {
	/* Initialize ThreadService */
	threadService := ThreadService{
		ctx:     ctx,
		queries: queries,
	}
	return &threadService
}

func (t *ThreadService) CreateThread(threadCreate schema.CreateThreadRequest) (domain.Thread, error) {
	/* Create a new thread */
	dbThread, err := t.queries.CreateThread(t.ctx, threadCreate.Title)
	domainThread := castRepositoryThreadToDomainThread(dbThread)
	return domainThread, err
}

func (t *ThreadService) GetThreadById(threadId uuid.UUID) (domain.Thread, error) {
	/* Retrieve a thread by its id. */
	dbThread, err := t.queries.GetThreadById(t.ctx, threadId)
	domainThread := castRepositoryThreadToDomainThread(dbThread)
	return domainThread, err
}

func (t *ThreadService) ListThreads() []domain.Thread {
	/* Get information of all threads - paginated */
	dbThreads, err := t.queries.ListThreads(t.ctx)
	if err != nil {
		return []domain.Thread{}
	}

	// Cast the repository.Thread s to the domain.Thread s
	var threads []domain.Thread
	for _, dbThread := range dbThreads {
		threads = append(threads, castRepositoryThreadToDomainThread(dbThread))
		return threads
	}
	return threads
}

func castRepositoryThreadToDomainThread(dbThread repository.Thread) domain.Thread {
	/* Helper function: Cast the database object to the domain object */

	domainThread := domain.Thread{
		ID:    dbThread.ID.String(),
		Title: dbThread.Title,

		CreatedAt: dbThread.CreatedAt,
		UpdatedAt: dbThread.UpdatedAt,
		DeletedAt: dbThread.DeletedAt,
	}
	return domainThread
}
