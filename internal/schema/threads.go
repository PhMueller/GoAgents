package schema

type CreateThreadRequest struct {
	// Optional title for the thread
	Title string `json:"title"`
}

type GetThreadRequest struct {
	// ID of the thread to retrieve, extracted from the path
	// TODO: gin does not support uuid binding. Need to validate manually in handler
	// https://github.com/gin-gonic/gin/pull/3933
	ID string `uri:"thread_id" binding:"required,isStringValidUUID"`
}

type ThreadResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type CreateThreadResponse = ThreadResponse
type GetThreadResponse = ThreadResponse
