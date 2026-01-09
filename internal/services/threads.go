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

func (t *ThreadService) GetThreadByID(threadID uuid.UUID) (domain.Thread, error) {
	/* Retrieve a thread by its id. */
	dbThread, err := t.queries.GetThreadById(t.ctx, threadID)
	domainThread := castRepositoryThreadToDomainThread(dbThread)
	return domainThread, err
}

func (t *ThreadService) GetThreadsInfo() []domain.Thread {
	/* Get information of all threads - paginated

	# TODO:
	- add pagination parameters
	- add filter on user
	*/
	dbThreads, err := t.queries.ListThreads(t.ctx)
	if err != nil {
		return []domain.Thread{}
	}

	// Cast the repository.Thread s to the domain.Thread s
	threads := make([]domain.Thread, len(dbThreads))
	for i, dbThread := range dbThreads {
		threads[i] = castRepositoryThreadToDomainThread(dbThread)
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
