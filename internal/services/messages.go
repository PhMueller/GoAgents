package services

import (
	"context"
	"log"

	"examples.com/assistants/db/repository"
	"examples.com/assistants/internal/schema"
	"github.com/google/uuid"
)

type MessageService struct {
	ctx     context.Context
	queries repository.Queries
}

func NewMessageService(ctx context.Context, queries repository.Queries) *MessageService {
	messageService := MessageService{
		ctx:     ctx,
		queries: queries,
	}
	log.Println("MessageService initialized")
	return &messageService
}

func (m *MessageService) GetMessageByMessageId(messageId uuid.UUID) repository.Message {
	message, err := m.queries.GetMessageByMessageId(m.ctx, messageId)
	if err != nil {
		// TODO: raise proper error
		return repository.Message{}
	}
	return message
}

func (m *MessageService) GetMessagesByThreadId(threadId uuid.UUID) []repository.Message {
	messages, err := m.queries.GetMessagesByThreadId(m.ctx, threadId)
	if err != nil {
		// TODO: raise proper error
		log.Println(err)
		return []repository.Message{}
	}
	if messages == nil {
		messages = []repository.Message{}
	}
	return messages
}

func (m *MessageService) CreateMessage(message schema.CreateMessageRequest) (repository.Message, error) {

	threadID := uuid.Must(uuid.Parse(message.ThreadId))

	createMessageParams := repository.CreateMessageParams{
		ThreadID: threadID,
		Content:  message.Content,
	}

	dbMessage, err := m.queries.CreateMessage(m.ctx, createMessageParams)
	return dbMessage, err
}
