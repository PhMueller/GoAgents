package routers

import (
	"examples.com/assistants/internal/services"
	"github.com/gin-gonic/gin"
)

type ClientSessionHandler struct {
	AuthService *services.AuthService
}

func NewClientSessionHandler(authService *services.AuthService) *ClientSessionHandler {
	return &ClientSessionHandler{AuthService: authService}
}

func (a *ClientSessionHandler) CreateClientSession(context *gin.Context) {}
func (a *ClientSessionHandler) GetClientSessions(context *gin.Context)   {}
func (a *ClientSessionHandler) GetClientSession(context *gin.Context)    {}
func (a *ClientSessionHandler) DeleteClientSession(context *gin.Context) {}
