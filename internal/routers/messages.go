package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"examples.com/assistants/internal/schema"
	"examples.com/assistants/internal/services"
)

type MessagesHandler struct {
	MessageService *services.MessageService
}

func NewMessagesHandler(messageService *services.MessageService) *MessagesHandler {
	return &MessagesHandler{MessageService: messageService}
}

func (m *MessagesHandler) GetMessagesByThreadId(w http.ResponseWriter, r *http.Request) {

	var msg schema.MessagesByThreadRead

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	messages := m.MessageService.GetMessagesByThreadId(msg.ThreadId)
	// TODO: How to error handling really!? :D

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "DecodeError", http.StatusInternalServerError)
	}
}

func (m *MessagesHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	threadIDStr := chi.URLParam(r, "thread_id")
	threadID, err := uuid.Parse(threadIDStr)
	if err != nil {
		InvalidThreadIdError(w, r)
		return
	}

	var msg schema.MessageCreate
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		InvalidRequestBodyError(w, r)
		return
	}

	createdMsg, err := m.MessageService.CreateMessage(threadID, msg)
	if err != nil {
		InternalServiceError(w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdMsg); err != nil {
		InternalServiceError(w, r)
		return
	}
}
