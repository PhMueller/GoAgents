package routers

import (
	"net/http"

	"examples.com/assistants/internal/schema"
	"examples.com/assistants/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ThreadsHandler struct {
	ThreadService *services.ThreadService
}

func NewThreadsHandler(threadService *services.ThreadService) *ThreadsHandler {
	return &ThreadsHandler{ThreadService: threadService}
}

func (t *ThreadsHandler) CreateThread(context *gin.Context) {
	var createThreadRequest schema.CreateThreadRequest
	if err := context.ShouldBind(&createThreadRequest); err != nil {
		context.JSON(GinInvalidRequestBodyError())
		return
	}

	createdThread, err := t.ThreadService.CreateThread(createThreadRequest)
	if err != nil {
		context.JSON(GinInternalServiceError())
		return
	}

	// Cast the repository object to schema.ThreadResponse
	title := ""
	if createdThread.Title != nil {
		title = *createdThread.Title
	}

	threadResponse := schema.ThreadResponse{
		ID:    createdThread.ID.String(),
		Title: title,
	}

	context.JSON(http.StatusOK, threadResponse)

}

func (t *ThreadsHandler) GetThreadById(context *gin.Context) {

	var request schema.GetThreadRequest

	if err := context.ShouldBindUri(&request); err != nil {
		//err, ok := err.(validator.ValidationErrors)
		//if ok {
		//	if err[0].Field() == "ID" {
		//		context.JSON(GinInvalidThreadIdError())
		//		return
		//	}
		//} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
		//}
	}

	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	threadId := uuid.Must(uuid.Parse(request.ID))

	thread, err := t.ThreadService.GetThreadById(threadId)
	if err != nil {
		context.JSON(GinInternalServiceError())
		return
	}

	title := ""
	if thread.Title != nil {
		title = *thread.Title
	}

	threadResponse := schema.ThreadResponse{
		ID:    thread.ID.String(),
		Title: title,
	}

	context.JSON(http.StatusOK, threadResponse)
}
