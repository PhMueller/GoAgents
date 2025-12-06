package routers

import (
	"encoding/json"
	"net/http"

	"examples.com/assistants/internal/schema"
	"examples.com/assistants/internal/services"
)

type MessagesHandler struct {
	MessageService *services.MessageService
}

func NewMessagesHandler(messageService *services.MessageService) *MessagesHandler {
	return &MessagesHandler{MessageService: messageService}
}

func (m *MessagesHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	messages := m.MessageService.GetMessages()
	// TODO: How to error handling :D

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "DecodeError", http.StatusInternalServerError)
	}
}

func (m *MessagesHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg schema.MessageCreate

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdMsg, err := m.MessageService.CreateMessage(msg)
	if err != nil {
		http.Error(w, "Failed to create message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdMsg); err != nil {
		http.Error(w, "Encode error", http.StatusInternalServerError)
	}
}
