package services

import (
	"context"
	"testing"
	"time"

	"examples.com/assistants/db/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

/*First create a mock for the messages queries */

type MockMessagesQueries struct {
	ReturnValue repository.Message
	ReturnError error
}

func (m *MockMessagesQueries) CreateMessage(ctx context.Context, arg repository.CreateMessageParams) (repository.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockMessagesQueries) CreateThread(ctx context.Context, title *string) (repository.Thread, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockMessagesQueries) DeleteMessage(ctx context.Context, id uuid.UUID) (repository.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockMessagesQueries) DeleteThread(ctx context.Context, id uuid.UUID) (repository.Thread, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockMessagesQueries) GetMessage(ctx context.Context, id uuid.UUID) (repository.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockMessagesQueries) GetMessages(ctx context.Context, id uuid.UUID) ([]repository.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockMessagesQueries) GetMessagesByThreadId(ctx context.Context, threadID uuid.UUID) ([]repository.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockMessagesQueries) GetThreadById(ctx context.Context, id uuid.UUID) (repository.Thread, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockMessagesQueries) ListThreads(ctx context.Context) ([]repository.Thread, error) {
	//TODO implement me
	panic("implement me")
}

func NewMockMessagesQueries(returnValue repository.Message, returnError error) *MockMessagesQueries {
	return &MockMessagesQueries{ReturnValue: returnValue, ReturnError: returnError}
}

func (m *MockMessagesQueries) GetMessageByMessageId(ctx context.Context, messageID uuid.UUID) (repository.Message, error) {
	return m.ReturnValue, m.ReturnError
}

func TestGetMessageByMessageID(t *testing.T) {
	// Arrange
	mockMessage := repository.Message{
		ID:        uuid.New(),
		ThreadID:  uuid.New(),
		Content:   "Test message",
		CreatedAt: time.Now(),
	}
	mockQueries := NewMockMessagesQueries(mockMessage, nil)

	// Act
	messageService := &MessageService{
		ctx:     context.Background(),
		queries: mockQueries,
	}

	domainMessage, err := messageService.GetMessageByMessageID(mockMessage.ID, mockMessage.ThreadID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if domainMessage.ID != mockMessage.ID {
		t.Errorf("Expected message ID %v, got %v", mockMessage.ID, domainMessage.ID)
	}

	if domainMessage.ThreadID != mockMessage.ThreadID {
		t.Errorf("Expected ThreadID %v, got %v", mockMessage.ThreadID, domainMessage.ThreadID)
	}

	if domainMessage.Content != mockMessage.Content {
		t.Errorf("Expected Content %v, got %v", mockMessage.Content, domainMessage.Content)
	}

	if domainMessage.CreatedAt != mockMessage.CreatedAt {
		t.Errorf("Expected CreatedAt %v, got %v", mockMessage.CreatedAt, domainMessage.CreatedAt)
	}

	assert.Nil(t, domainMessage.UpdatedAt)
	assert.Nil(t, domainMessage.DeletedAt)

}
