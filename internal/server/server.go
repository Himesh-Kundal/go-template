package server

import (
	"net/http"

	"go-template/internal/db"
	"go-template/internal/handlers"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router  *chi.Mux
	handler *handlers.Handler
}

func New(q *db.Queries) *Server {
	h := handlers.New(q)
	s := &Server{
		router:  chi.NewRouter(),
		handler: h,
	}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
