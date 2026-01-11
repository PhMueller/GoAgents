package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"examples.com/assistants/db/repository"
	"examples.com/assistants/internal/domain"
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
	/* Test the messageService's GetMessageByMessageID method */
	testCases := []struct {
		name        string
		mockMessage repository.Message
		mockError   error
	}{
		{
			name: "Successful retrieval of message by MessageID",
			mockMessage: repository.Message{
				ID:        uuid.New(),
				ThreadID:  uuid.New(),
				Content:   "Test message",
				CreatedAt: time.Now(),
			},
			mockError: nil,
		},
		{
			name:        "No message found for given MessageID",
			mockMessage: repository.Message{},
			mockError:   errors.New("message not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockContext := context.Background()

			var expectedMessage domain.Message
			if tc.mockError == nil {
				expectedMessage = domain.Message{
					ID:        tc.mockMessage.ID,
					ThreadID:  tc.mockMessage.ThreadID,
					Content:   tc.mockMessage.Content,
					CreatedAt: tc.mockMessage.CreatedAt,
				}
			} else {
				expectedMessage = domain.Message{}
			}

			mockQueries := NewMockMessagesQueries(tc.mockMessage, tc.mockError)

			// Act
			messageService := &MessageService{
				queries: mockQueries,
			}

			domainMessage, err := messageService.GetMessageByMessageID(
				mockContext, tc.mockMessage.ID, tc.mockMessage.ThreadID,
			)

			// Assert
			if tc.mockError == nil && err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			assert.Equal(t, expectedMessage, domainMessage)
			assert.Equal(t, tc.mockError, err)

		})
	}

}
