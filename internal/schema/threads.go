package schema

import "time"

type CreateThreadRequest struct {
	/* Input object for the POST /threads endpoint */
	Title *string `json:"title"`
}

type GetThreadRequest struct {
	/* Input object for the GET /threads/:thread_id endpoint */

	// ID of the thread to retrieve, extracted from the path
	// TODO: gin does not support uuid binding. Need to validate manually in handler
	// https://github.com/gin-gonic/gin/pull/3933
	// TODO: how to test the binding validation? we cannot use validator in the test with binding!
	ID string `uri:"thread_id" validate:"required,isStringValidUUID" binding:"required,isStringValidUUID"`
}

type GetThreadsInfoRequest struct {
	/* Input object for the GET /threads endpoint

	Supports optionally pagination via cursor and size parameters.
	*/
	Cursor *string `form:"cursor"`
	Size   *int    `form:"size"`
}

type ThreadResponse struct {
	/* Base Response object for thread-related endpoints */
	ID    string  `json:"id"`
	Title *string `json:"title"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CreateThreadResponse = ThreadResponse
type GetThreadResponse = ThreadResponse

type GetThreadsInfoResponse struct {
	/* Response object for the GET /threads endpoint.

	This object contains a list of threads and a cursor for pagination.
	*/
	Threads []ThreadResponse `json:"threads"`
	Cursor  *string          `json:"cursor,omitempty"`
	Size    *int             `json:"size,omitempty"`
}
