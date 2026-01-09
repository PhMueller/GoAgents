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

	domainThread, err := t.ThreadService.CreateThread(createThreadRequest)
	if err != nil {
		context.JSON(GinInternalServiceError())
		return
	}

	threadResponse := schema.CreateThreadResponse{
		ID:    domainThread.ID,
		Title: domainThread.Title,
	}

	context.JSON(http.StatusOK, threadResponse)

}

func (t *ThreadsHandler) GetThreadById(context *gin.Context) {

	var getThreadRequest schema.GetThreadRequest

	// We use the validator package to check for errors in the path parameters. They can be handled separately.
	if err := context.ShouldBindUri(&getThreadRequest); err != nil {
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

	if err := context.ShouldBind(&getThreadRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We have validated that the id is a valid uuid in the schema binding step, so we can safely parse it here.
	threadId := uuid.Must(uuid.Parse(getThreadRequest.ID))

	domainThread, err := t.ThreadService.GetThreadById(threadId)
	if err != nil {
		context.JSON(GinInternalServiceError())
		return
	}

	threadResponse := schema.GetThreadResponse{
		ID:    domainThread.ID,
		Title: domainThread.Title,
	}

	context.JSON(http.StatusOK, threadResponse)
}
