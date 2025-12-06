package repositories

import (
	"examples.com/assistants/internal/domain"
	"examples.com/assistants/internal/schema"
	"github.com/google/uuid"

	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type MessageRepository struct {
	db *sqlx.DB
}

const (
	MessageTableName = "messages"
)

func NewMessageRepository(db *sqlx.DB) *MessageRepository {
	messageRepository := MessageRepository{db: db}
	return &messageRepository
}

func (repo *MessageRepository) Create(message schema.MessageCreate) (domain.Message, error) {
	query := fmt.Sprintf(`INSERT INTO %s (thread_id, content) VALUES ($1, $2) RETURNING id`, MessageTableName)

	var id uuid.UUID
	err := repo.db.QueryRow(query, message.ThreadId, message.Content).Scan(&id)
	if err != nil {
		return domain.Message{}, err
	}

	domainMessage := domain.Message{
		ID:       id,
		ThreadId: message.ThreadId,
		Content:  message.Content,
	}
	return domainMessage, nil
}

func (repo *MessageRepository) Get(id uuid.UUID) (domain.Message, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = ($1)`, MessageTableName)
	var message domain.Message
	err := repo.db.Get(&message, query, id)
	return message, err
}

func (repo *MessageRepository) List() ([]domain.Message, error) {
	query := fmt.Sprintf(`SELECT * FROM %s`, MessageTableName)
	var messages []domain.Message
	err := repo.db.Select(&messages, query)
	return messages, err
}
