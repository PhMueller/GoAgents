package routers

import (
	"net/http"

	"examples.com/assistants/internal/schema"
	"examples.com/assistants/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MessagesHandler struct {
	MessageService *services.MessageService
}

func NewMessagesHandler(messageService *services.MessageService) *MessagesHandler {
	return &MessagesHandler{MessageService: messageService}
}

func (m *MessagesHandler) GetMessagesByThreadId(context *gin.Context) {

	var request schema.GetMessageRequest

	err := context.ShouldBindUri(&request)
	if err != nil {
		context.JSON(GinInvalidThreadIdError())
		return
	}

	err = context.ShouldBind(&request)
	if err != nil {
		context.JSON(GinInvalidThreadIdError())
		return
	}

	threadId := uuid.Must(uuid.Parse(request.ThreadId))

	// TODO: How to do proper error handling?
	messages := m.MessageService.GetMessagesByThreadId(threadId)

	messagesRead := make([]schema.GetMessageResponse, len(messages))
	for i, message := range messages {
		messagesRead[i] = schema.GetMessageResponse{
			ID:       message.ID,
			ThreadId: message.ThreadID,
			Content:  message.Content,
		}
	}
	context.JSON(http.StatusOK, messagesRead)
}

func (m *MessagesHandler) CreateMessage(context *gin.Context) {

	var request schema.CreateMessageRequest
	err := context.ShouldBindUri(&request)
	if err != nil {
		context.JSON(GinInvalidThreadIdError())
		return
	}

	threadID := uuid.Must(uuid.Parse(request.ThreadId))

	err = context.ShouldBind(&request)
	if err != nil {
		// TODO: decide if the thread id is missing or the body is invalid.
		context.JSON(GinInvalidThreadIdError())
		return
	}

	createdMsg, err := m.MessageService.CreateMessage(request)
	if err != nil {
		context.JSON(GinInternalServiceError())
		return
	}

	createMessageResponse := schema.CreateMessageResponse{
		ID:       createdMsg.ID,
		ThreadId: threadID,
		Content:  createdMsg.Content,
	}
	//responseJson, err := json.Marshal(createMessageResponse)
	//if err != nil {
	//	context.JSON(GinInternalServiceError())
	//	return
	//}

	context.JSON(http.StatusCreated, createMessageResponse)
}
