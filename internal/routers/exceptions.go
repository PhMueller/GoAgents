package routers

import "net/http"

func InvalidThreadIdError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid thread id", http.StatusBadRequest)
}

func InvalidRequestBodyError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid request body", http.StatusBadRequest)
}

func InternalServiceError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal service error", http.StatusInternalServerError)
}
