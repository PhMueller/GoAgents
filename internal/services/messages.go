package services

import (
	"context"
	"log"

	"examples.com/assistants/db/repository"
	"examples.com/assistants/internal/domain"
	"examples.com/assistants/internal/schema"
	"github.com/google/uuid"
)

type MessageService struct {
	ctx     context.Context
	queries repository.Querier
}

func NewMessageService(ctx context.Context, queries repository.Querier) *MessageService {
	/* Initialize MessageService */
	messageService := MessageService{
		ctx:     ctx,
		queries: queries,
	}
	log.Println("MessageService initialized")
	return &messageService
}

func (m *MessageService) GetMessageByMessageID(messageID uuid.UUID, threadID uuid.UUID) (domain.Message, error) {
	/* Retrieve a message by its id */

	// TODO: use the thread id to verify that the message belongs to the thread and the user has access to the thread.

	dbMessage, err := m.queries.GetMessageByMessageId(m.ctx, messageID)
	if err != nil {
		// TODO: raise proper error
		return domain.Message{}, err
	}
	domainMessage := castRepositoryMessageToDomainMessage(dbMessage)

	return domainMessage, nil
}

func (m *MessageService) GetMessagesByThreadID(threadID uuid.UUID) []domain.Message {
	/* Retrieve all messages in a thread */
	dbMessages, err := m.queries.GetMessagesByThreadId(m.ctx, threadID)
	if err != nil {
		// TODO: raise proper error
		log.Println(err)
		return []domain.Message{}
	}
	if dbMessages == nil {
		return []domain.Message{}
	}

	domainMessages := make([]domain.Message, len(dbMessages))
	for i, msg := range dbMessages {
		domainMessages[i] = castRepositoryMessageToDomainMessage(msg)
	}

	return domainMessages
}

func (m *MessageService) CreateMessage(message schema.CreateMessageRequest) (domain.Message, error) {
	/* Create a new message in a thread */
	threadID := uuid.Must(uuid.Parse(message.ThreadID))

	createMessageParams := repository.CreateMessageParams{
		ThreadID: threadID,
		Content:  message.Content,
	}

	dbMessage, err := m.queries.CreateMessage(m.ctx, createMessageParams)
	domainMessage := castRepositoryMessageToDomainMessage(dbMessage)
	return domainMessage, err
}

func castRepositoryMessageToDomainMessage(dbMessage repository.Message) domain.Message {
	/* Helper function: Cast the database object to the domain object */
	domainMessage := domain.Message{
		ID:       dbMessage.ID,
		ThreadID: dbMessage.ThreadID,
		Content:  dbMessage.Content,

		CreatedAt: dbMessage.CreatedAt,
		UpdatedAt: dbMessage.UpdatedAt,
		DeletedAt: dbMessage.DeletedAt,
	}
	return domainMessage
}
