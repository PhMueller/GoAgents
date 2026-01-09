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

func (t *ThreadsHandler) GetThreadByID(context *gin.Context) {

	var getThreadRequest schema.GetThreadRequest

	// We use the validator package to check for errors in the path parameters. They can be handled separately.
	if err := context.ShouldBindUri(&getThreadRequest); err != nil {
		//err, ok := err.(validator.ValidationErrors)
		//if ok {
		//	if err[0].Field() == "ID" {
		//		context.JSON(GinInvalidThreadIDError())
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
	threadID := uuid.Must(uuid.Parse(getThreadRequest.ID))

	domainThread, err := t.ThreadService.GetThreadByID(threadID)
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

func (t *ThreadsHandler) ListThreads(context *gin.Context) {
	/* This route returns the list of thread ids that are available for a user

	Notes:
	- This endpoint supports cursor based pagination.
	- The `cursor` field can be used to fetch the next page of results.
	- If there are no more results, the `cursor` field will be omitted.
	*/
	var request schema.GetThreadsRequest
	if err := context.ShouldBindQuery(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
