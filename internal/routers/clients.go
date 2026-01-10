package routers

import (
	"examples.com/assistants/internal/services"
	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	AuthService *services.AuthService
}

func NewClientHandler(authService *services.AuthService) *ClientHandler {
	return &ClientHandler{AuthService: authService}
}

func (a *ClientHandler) CreateClient(context *gin.Context) {}
func (a *ClientHandler) GetClients(context *gin.Context)   {}
func (a *ClientHandler) GetClient(context *gin.Context)    {}
func (a *ClientHandler) UpdateClient(context *gin.Context) {}
func (a *ClientHandler) DeleteClient(context *gin.Context) {}
