package server

import (
	"github.com/go-chi/chi/v5/middleware"

	internalMiddleware "go-template/internal/middleware"
)

func (s *Server) routes() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.router.Post("/auth/register", s.handler.HandleCreateUser)
	s.router.Post("/auth/login", s.handler.HandleLogin)

	s.router.With(internalMiddleware.AuthMiddleware).Get("/users", s.handler.HandleGetUserByEmail)
}
