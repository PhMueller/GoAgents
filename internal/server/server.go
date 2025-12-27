package server

import (
	"examples.com/assistants/internal/routers"
	"examples.com/assistants/internal/schema"
	"examples.com/assistants/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	server := &Server{Engine: gin.Default()}

	return server
}

func (s *Server) AddValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("isStringValidUUID", schema.IsStringValidUUID)
	}
}

func (s *Server) SetupRoutes(messagesService *services.MessageService, threadService *services.ThreadService) {
	messagesHandler := routers.NewMessagesHandler(messagesService)
	threadsHandler := routers.NewThreadsHandler(threadService)

	v1 := s.Engine.Group("/v1")
	messagesRouter := v1.Group("/messages")
	messagesRouter.GET("/:thread_id", messagesHandler.GetMessagesByThreadId)
	messagesRouter.POST("/messages/:thread_id", messagesHandler.CreateMessage)

	threadsRouter := v1.Group("/threads")
	threadsRouter.POST("", threadsHandler.CreateThread)
	threadsRouter.GET("/:thread_id", threadsHandler.GetThreadById)
}
