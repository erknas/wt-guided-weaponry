package api

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/zeze322/wt-guided-weaponry/lib"
)

func (s *Server) handleCategories(w http.ResponseWriter, r *http.Request) error {
	categories, err := s.postgres.Categories(r.Context())
	if err != nil {
		return err
	}

	s.logger.WithFields(log.Fields{
		"request_id": fmt.Sprintf("%d", os.Getpid()),
		"method":     r.Method,
		"path":       r.URL.Path,
	}).Info("got categories", categories)

	return lib.WriteJSON(w, http.StatusOK, categories)
}
