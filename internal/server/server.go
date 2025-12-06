package server

import (
	"examples.com/assistants/internal/routers"
	"examples.com/assistants/internal/services"
	"github.com/go-chi/chi/v5"
)

//// ... router/health.go

//type HealthResponse struct {
//	Status string `json:"status"`
//}
//
//// Implement a health endpoint that returns 200 if it is alive.
//func health(writer http.ResponseWriter, request *http.Request) {
//
//	response := HealthResponse{Status: "ok"}
//	writer.WriteHeader(http.StatusOK)
//
//	if err := json.NewEncoder(writer).Encode(response); err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}

// ../server

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	server := &Server{Router: chi.NewRouter()}
	return server
}

func (s *Server) SetupRoutes(messagesService *services.MessageService) {
	messagesHandler := routers.NewMessagesHandler(messagesService)
	s.Router.Get("/messages", messagesHandler.GetMessagesByThreadId)
	s.Router.Post("/messages/create", messagesHandler.CreateMessage)
}
