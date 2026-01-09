package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinInvalidThreadIDError() (int, gin.H) {
	return http.StatusBadRequest, gin.H{"error": "Invalid thread_id"}
}

func GinInvalidRequestBodyError() (int, gin.H) {
	return http.StatusBadRequest, gin.H{"error": "Invalid request body"}
}

func GinInternalServiceError() (int, gin.H) {
	return http.StatusInternalServerError, gin.H{"error": "Internal service error"}
}
