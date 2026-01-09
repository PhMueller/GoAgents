package schema

import (
	"time"

	"github.com/google/uuid"
)

type CreateMessageRequest struct {
	/* Input object for the POST /threads/:thread_id/messages endpoint */

	// ThreadID comes from the path parameter
	ThreadID string `form:"thread_id" binding:"required,isStringValidUUID"`
	Content  string `json:"content" binding:"required"`
}

type GetMessageRequest struct {
	/* Input object for the GET /threads/:thread_id/messages/:message_id endpoint */

	// ThreadID comes from the path parameter
	ThreadID  string `form:"thread_id" binding:"required,isStringValidUUID"`
	MessageID string `form:"message_id" binding:"required,isStringValidUUID"`
}

type GetMessagesRequest struct {
	/* Input object for the GET /threads/:thread_id/messages endpoint */

	// ThreadID comes from the path parameter
	ThreadID string `form:"thread_id" binding:"required,isStringValidUUID"`
}

type MessageResponse struct {
	/* Base Response object for message-related endpoints */
	ID       uuid.UUID `json:"id"`
	ThreadID uuid.UUID `json:"thread_id"`
	Content  string    `json:"content"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CreateMessageResponse = MessageResponse
type GetMessageResponse = MessageResponse

type GetMessagesResponse struct {
	/* Response object for the GET /threads/:thread_id/messages endpoint
	Optionally paginated
	*/
	Messages []MessageResponse `json:"messages"`
	Cursor   *string           `json:"cursor,omitempty"`
	Size     *int              `json:"size,omitempty"`
}
