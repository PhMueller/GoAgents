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

func (m *MessagesHandler) CreateMessage(context *gin.Context) {
	/* Create a new message in a thread */
	var request schema.CreateMessageRequest

	// We ignore the first command since it always returns an error (the id is not present yet)
	// Checking the error in the second step is sufficient.
	// https://github.com/gin-gonic/gin/issues/2758
	err := context.ShouldBind(&request)
	err = context.ShouldBindUri(&request)
	if err != nil {
		//	// TODO: Add fine grained error handling.
		//	//       - Raise an error if the thread id is missing in the URI.
		//	//       - Raise an error if the body is invalid.
		//	context.JSON(GinInvalidThreadIDError())
		context.JSON(GinInvalidThreadIDError())
		return
	}

	threadID := uuid.Must(uuid.Parse(request.ThreadID))

	domainMessage, err := m.MessageService.CreateMessage(request)
	if err != nil {
		context.JSON(GinInternalServiceError())
		return
	}

	responseMessage := schema.CreateMessageResponse{
		ID:        domainMessage.ID,
		ThreadID:  threadID,
		Content:   domainMessage.Content,
		CreatedAt: domainMessage.CreatedAt,
		UpdatedAt: domainMessage.UpdatedAt,
		DeletedAt: domainMessage.DeletedAt,
	}

	context.JSON(http.StatusCreated, responseMessage)
}

func (m *MessagesHandler) GetMessageByMessageID(context *gin.Context) {
	/* Retrieve a message by its id */
	var request schema.GetMessageRequest

	err := context.ShouldBindUri(&request)
	if err != nil {
		// TODO: better handling. Either MessageID invalid or ThreadID invalid
		context.JSON(GinInternalServiceError())
		return
	}

	// We have validated the UUIDs in the binding step. Thus, we can safely parse them.
	messageID := uuid.Must(uuid.Parse(request.MessageID))
	threadID := uuid.Must(uuid.Parse(request.ThreadID))

	domainMessage, err := m.MessageService.GetMessageByMessageID(messageID, threadID)
	if err != nil {
		context.JSON(GinInternalServiceError())
		return
	}

	messageResponse := schema.GetMessageResponse{
		ID:        domainMessage.ID,
		ThreadID:  domainMessage.ThreadID,
		Content:   domainMessage.Content,
		CreatedAt: domainMessage.CreatedAt,
		UpdatedAt: domainMessage.UpdatedAt,
		DeletedAt: domainMessage.DeletedAt,
	}

	context.JSON(http.StatusOK, messageResponse)
}

func (m *MessagesHandler) GetMessagesByThreadID(context *gin.Context) {
	/* Retrieve all messages in a thread */
	var request schema.GetMessagesRequest

	err := context.ShouldBindUri(&request)
	if err != nil {
		context.JSON(GinInvalidThreadIDError())
		return
	}

	threadID := uuid.Must(uuid.Parse(request.ThreadID))

	// TODO: How to do proper error handling?
	domainMessages := m.MessageService.GetMessagesByThreadID(threadID)

	messageResponseItems := make([]schema.GetMessageResponse, len(domainMessages))
	for i, message := range domainMessages {
		messageResponseItems[i] = schema.GetMessageResponse{
			ID:        message.ID,
			ThreadID:  message.ThreadID,
			Content:   message.Content,
			CreatedAt: message.CreatedAt,
			UpdatedAt: message.UpdatedAt,
			DeletedAt: message.DeletedAt,
		}
	}

	messagesResponse := schema.GetMessagesResponse{
		Messages: messageResponseItems,
		Cursor:   nil,
		Size:     nil,
	}

	context.JSON(http.StatusOK, messagesResponse)
}
