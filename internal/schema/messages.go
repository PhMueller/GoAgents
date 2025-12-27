package schema

import (
	"github.com/google/uuid"
)

type GetMessageRequest struct {
	// the thread id comes from the path parameter
	ThreadId string `form:"thread_id" binding:"required,isStringValidUUID"`
}

type CreateMessageRequest struct {
	// the thread id comes from the path parameter
	ThreadId string `form:"thread_id" binding:"required,isStringValidUUID"`
	Content  string `json:"content" binding:"required"`
}

type MessageResponse struct {
	ID       uuid.UUID `json:"id"`
	ThreadId uuid.UUID `json:"thread_id"`
	Content  string    `json:"content"`
}

type CreateMessageResponse = MessageResponse
type GetMessageResponse = MessageResponse
