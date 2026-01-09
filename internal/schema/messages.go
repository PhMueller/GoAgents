package schema

import (
	"github.com/google/uuid"
)

type CreateMessageRequest struct {
	/* Input object for the POST /threads/:thread_id/messages endpoint */

	// ThreadId comes from the path parameter
	ThreadId string `form:"thread_id" binding:"required,isStringValidUUID"`
	Content  string `json:"content" binding:"required"`
}

type GetMessageRequest struct {
	/* Input object for the GET /threads/:thread_id/messages/:message_id endpoint */

	// ThreadId comes from the path parameter
	ThreadId string `form:"thread_id" binding:"required,isStringValidUUID"`
}

type MessageResponse struct {
	/* Base Response object for message-related endpoints */
	ID       uuid.UUID `json:"id"`
	ThreadId uuid.UUID `json:"thread_id"`
	Content  string    `json:"content"`
}

type CreateMessageResponse = MessageResponse
type GetMessageResponse = MessageResponse
