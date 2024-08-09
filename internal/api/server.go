package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"

	"github.com/zeze322/wt-guided-weaponry/internal/db/postgresdb"
	"github.com/zeze322/wt-guided-weaponry/lib"
)

type Server struct {
	logger   *log.Logger
	port     string
	postgres postgresdb.Storage
}

func NewServer(logger *log.Logger, port string, postgres postgresdb.Storage) *Server {
	return &Server{
		logger:   logger,
		port:     port,
		postgres: postgres,
	}
}

func (s *Server) Run() error {
	router := chi.NewRouter()

	router.Get("/categories", lib.MakeHTTP(s.handleCategories))

	log.Printf("Starting server on http://localhost:8000")

	if err := http.ListenAndServe(s.port, router); err != nil {
		return fmt.Errorf("failed to start server: %s", err)
	}

	return nil
}
