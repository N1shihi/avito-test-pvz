package server

import (
	"avito-test-pvz/internal/config"
	"avito-test-pvz/internal/handler"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	cfg    *config.Config
	db     *sql.DB
}

func New(cfg *config.Config, db *sql.DB) *Server {

	return &Server{
		router: gin.Default(),
		cfg:    cfg,
		db:     db,
	}
}
func (s *Server) Run(h *handler.Handler) {
	s.MapRoutes(h)
	s.router.Run(s.cfg.Server.Address)
}
