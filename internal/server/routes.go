package server

import "avito-test-pvz/internal/handler"

func (s *Server) MapRoutes(h *handler.Handler) {
	s.router.POST("/register", h.CreateUser)
}
