package routers

import (
	"examples.com/assistants/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (a *AuthHandler) CreateToken(context *gin.Context) {}
