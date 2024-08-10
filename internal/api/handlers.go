package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"

	"github.com/zeze322/wt-guided-weaponry/lib"
	"github.com/zeze322/wt-guided-weaponry/models"
)

type CategoriesResponse struct {
	Categories []models.Category `json:"categories"`
}

type WeaponsResponse struct {
	Weapons []*models.Params `json:"weapons"`
}

func (s *Server) handleWeapon(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleWeaponByName(w, r)
	case "PUT":
		return s.handleUpdateWeapon(w, r)
	case "DELETE":
		return s.handleDeleteWeapon(w, r)
	}

	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *Server) handleCategories(w http.ResponseWriter, r *http.Request) error {
	categories, err := s.postgres.Categories(r.Context())
	if err != nil {
		s.logger.WithFields(log.Fields{
			"request_id": fmt.Sprintf("%d", os.Getpid()),
			"path":       r.URL.Path,
			"error":      err,
		}).Debug("handleCategories error")
		return err
	}

	s.logger.Infof("got %d categories", len(categories))

	return lib.WriteJSON(w, http.StatusOK, CategoriesResponse{Categories: categories})
}

func (s *Server) handleWeapons(w http.ResponseWriter, r *http.Request) error {
	weapons, err := s.mongo.Weapons(r.Context())
	if err != nil {
		s.logger.WithFields(log.Fields{
			"request_id": fmt.Sprintf("%d", os.Getpid()),
			"path":       r.URL.Path,
			"error":      err,
		}).Debug("handleWeapons error")
		return err
	}

	s.logger.Infof("got %d weapons", len(weapons))

	return lib.WriteJSON(w, http.StatusOK, WeaponsResponse{Weapons: weapons})
}

func (s *Server) handleWeaponByName(w http.ResponseWriter, r *http.Request) error {
	name := chi.URLParam(r, "name")

	weapon, err := s.mongo.WeaponByName(r.Context(), name)
	if err != nil {
		s.logger.WithFields(log.Fields{
			"request_id": fmt.Sprintf("%d", os.Getpid()),
			"path":       r.URL.Path,
			"error":      err,
		}).Debug("handleWeaponByName error")
		return err
	}

	s.logger.Infof("got %s weapon", name)

	return lib.WriteJSON(w, http.StatusOK, weapon)
}

func (s *Server) handleWeaponsByCategory(w http.ResponseWriter, r *http.Request) error {
	category := chi.URLParam(r, "category")

	weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
	if err != nil {
		s.logger.WithFields(log.Fields{
			"request_id": fmt.Sprintf("%d", os.Getpid()),
			"path":       r.URL.Path,
			"error":      err,
		}).Debug("handleWeaponsByCategory error")
		return err
	}

	s.logger.Infof("got %d weapons", len(weapons))

	return lib.WriteJSON(w, http.StatusOK, WeaponsResponse{Weapons: weapons})
}

func (s *Server) handleInsertWeapon(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed: %s", r.Method)
	}

	req := new(models.Params)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	if err := s.mongo.InsertWeapon(r.Context(), req); err != nil {
		s.logger.WithFields(log.Fields{
			"request_id": fmt.Sprintf("%d", os.Getpid()),
			"path":       r.URL.Path,
			"error":      err,
		}).Debug("handleInsertWeapon error")
		return err
	}

	s.logger.Infof("insert %s weapon", req.Name)

	return lib.WriteJSON(w, http.StatusOK, struct{}{})
}

func (s *Server) handleUpdateWeapon(w http.ResponseWriter, r *http.Request) error {
	name := chi.URLParam(r, "name")

	req := new(models.Params)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.logger.WithFields(log.Fields{
			"request_id": fmt.Sprintf("%d", os.Getpid()),
			"path":       r.URL.Path,
			"error":      err,
		}).Debug("handleUpdateWeapon json decode error")
		return err
	}

	if err := s.mongo.UpdateWeapon(r.Context(), name, req); err != nil {
		s.logger.WithFields(log.Fields{
			"request_id": fmt.Sprintf("%d", os.Getpid()),
			"path":       r.URL.Path,
			"error":      err,
		}).Debug("handleUpdateWeapon error")
		return err
	}

	s.logger.Infof("update %s weapon", req.Name)

	return lib.WriteJSON(w, http.StatusOK, struct{}{})
}

func (s *Server) handleDeleteWeapon(w http.ResponseWriter, r *http.Request) error {
	name := chi.URLParam(r, "name")

	if err := s.mongo.DeleteWeapon(r.Context(), name); err != nil {
		s.logger.WithFields(log.Fields{
			"request_id": fmt.Sprintf("%d", os.Getpid()),
			"path":       r.URL.Path,
			"error":      err,
		}).Debug("handleDeleteWeapon error")
		return err
	}

	s.logger.Infof("delete %s weapon", name)

	return lib.WriteJSON(w, http.StatusOK, struct{}{})
}

// TODO handleSearchWeapon

// TODO handleCompareWeapons
