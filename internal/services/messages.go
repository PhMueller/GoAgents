package services

import (
	"log"

	"examples.com/assistants/internal/domain"
	"examples.com/assistants/internal/repositories"
	"examples.com/assistants/internal/schema"
)

type MessageService struct {
	messageRepository *repositories.MessageRepository
}

func NewMessageService(messageRepository *repositories.MessageRepository) *MessageService {
	messageService := MessageService{messageRepository: messageRepository}
	return &messageService
}

func (m *MessageService) GetMessages() []domain.Message {
	messages, err := m.messageRepository.List()
	if err != nil {
		return []domain.Message{}
	}
	return messages
}

func (m *MessageService) CreateMessage(message schema.MessageCreate) (domain.Message, error) {
	domainMessage, err := m.messageRepository.Create(message)
	return domainMessage, err
}
