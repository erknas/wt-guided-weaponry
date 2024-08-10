package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"

	"github.com/zeze322/wt-guided-weaponry/internal/db/mongodb"
	"github.com/zeze322/wt-guided-weaponry/internal/db/postgresdb"
	"github.com/zeze322/wt-guided-weaponry/lib"
)

type Server struct {
	logger   *log.Logger
	port     string
	postgres postgresdb.Store
	mongo    mongodb.Store
}

func NewServer(logger *log.Logger, port string, postgres postgresdb.Store, mongo mongodb.Store) *Server {
	return &Server{
		logger:   logger,
		port:     port,
		postgres: postgres,
		mongo:    mongo,
	}
}

func (s *Server) Run() error {
	router := chi.NewRouter()

	router.HandleFunc("/categories", lib.MakeHTTP(s.handleCategories))
	router.HandleFunc("/categories/{category}", lib.MakeHTTP(s.handleWeaponsByCategory))
	router.HandleFunc("/weapons", lib.MakeHTTP(s.handleWeapons))
	router.HandleFunc("/weapon", lib.MakeHTTP(s.handleInsertWeapon))
	router.HandleFunc("/weapon/{name}", lib.MakeHTTP(s.handleWeapon))

	log.Printf("Starting server on http://localhost:8000")

	if err := http.ListenAndServe(s.port, router); err != nil {
		return fmt.Errorf("failed to start server: %s", err)
	}

	return nil
}
