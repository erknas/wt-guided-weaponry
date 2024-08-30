package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/zeze322/wt-guided-weaponry/internal/db/mongodb"
	"github.com/zeze322/wt-guided-weaponry/lib"
)

type Server struct {
	port  string
	mongo mongodb.Store
}

func NewServer(port string, mongo mongodb.Store) *Server {
	return &Server{
		port:  port,
		mongo: mongo,
	}
}

func (s *Server) Run() error {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Handle("/*", public())

	router.Get("/", lib.MakeHTTP(s.handleHome))
	router.Get("/category", lib.MakeHTTP(s.handleCategories))
	router.Get("/category", lib.MakeHTTP(s.handleWeaponsByCategory))
	router.HandleFunc("/weapon/{name}", lib.MakeHTTP(s.handleWeapon))
	router.Post("/weapon", lib.MakeHTTP(s.handleInsertWeapon))
	router.Get("/weapons", lib.MakeHTTP(s.handleSearchWeapon))
	router.Get("/weapons", lib.MakeHTTP(s.handleWeapons))

	log.Printf("Running on http://localhost%s", s.port)

	if err := http.ListenAndServe(s.port, router); err != nil {
		return fmt.Errorf("failed to start server: %s", err)
	}

	return nil
}

func public() http.Handler {
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}
