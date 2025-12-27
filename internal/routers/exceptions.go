package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InvalidThreadIdError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid thread id", http.StatusBadRequest)
}

func GinInvalidThreadIdError() (int, gin.H) {
	return http.StatusBadRequest, gin.H{"error": "Invalid thread id"}
}

func InvalidRequestBodyError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid request body", http.StatusBadRequest)
}
func GinInvalidRequestBodyError() (int, gin.H) {
	return http.StatusBadRequest, gin.H{"error": "Invalid request body"}
}

func InternalServiceError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal service error", http.StatusInternalServerError)
}

func GinInternalServiceError() (int, gin.H) {
	return http.StatusInternalServerError, gin.H{"error": "Internal service error"}
}
