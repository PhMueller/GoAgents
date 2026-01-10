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
	/* Validators can be added to the validators engine (validators package)

	These are functions that are referenced in the `binding` tags in the schema definitions.

	https://github.com/go-playground/validator/blob/master/_examples/struct-level/main.go
	*/
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("isStringValidUUID", schema.IsStringValidUUID)
	}
}

func (s *Server) SetupRoutes(messagesService *services.MessageService, threadService *services.ThreadService, authService *services.AuthService) {
	messagesHandler := routers.NewMessagesHandler(messagesService)
	threadsHandler := routers.NewThreadsHandler(threadService)
	authHandler := routers.NewAuthHandler(authService)
	clientHandler := routers.NewClientHandler(authService)
	clientSessionHandler := routers.NewClientSessionHandler(authService)

	// Client: Chat interactions
	v1 := s.Engine.Group("/v1")

	threadsRouter := v1.Group("/threads")
	threadsRouter.POST("", threadsHandler.CreateThread)
	threadsRouter.GET("", threadsHandler.GetThreadsInfo)
	threadsRouter.GET("/:thread_id", threadsHandler.GetThreadByID)

	messagesRouter := v1.Group("/threads/:thread_id/messages")
	messagesRouter.POST("", messagesHandler.CreateMessage)
	messagesRouter.GET("", messagesHandler.GetMessagesByThreadID)
	messagesRouter.GET("/:message_id", messagesHandler.GetMessageByMessageID)

	// Client: Authentication routes
	authRouter := s.Engine.Group("/auth")
	authRouter.POST("/token", authHandler.CreateToken)

	// Admin Space: Client and session management
	adminRouter := s.Engine.Group("/admin")

	clientRouter := adminRouter.Group("/clients")
	clientRouter.POST("", clientHandler.CreateClient)
	clientRouter.GET("", clientHandler.GetClients)
	clientRouter.GET("/:client_id", clientHandler.GetClient)
	clientRouter.PATCH("/:client_id", clientHandler.UpdateClient)
	clientRouter.DELETE("/:client_id", clientHandler.DeleteClient)

	clientSessionRouter := clientRouter.Group("/:client_id/sessions")
	clientSessionRouter.POST("", clientSessionHandler.CreateClientSession)
	clientSessionRouter.GET("", clientSessionHandler.GetClientSessions)
	clientSessionRouter.GET("/:session_id", clientSessionHandler.GetClientSession)
	clientSessionRouter.DELETE("/:session_id", clientSessionHandler.DeleteClientSession)
}
