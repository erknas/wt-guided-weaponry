package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/zeze322/wt-guided-weaponry/lib"
	"github.com/zeze322/wt-guided-weaponry/models"
	"github.com/zeze322/wt-guided-weaponry/views/home"
	"github.com/zeze322/wt-guided-weaponry/views/rearaspect"
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
	}

	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *Server) handleCategories(w http.ResponseWriter, r *http.Request) error {
	_, err := s.mongo.Categories(r.Context())
	if err != nil {
		return err
	}

	return lib.Render(w, r, home.Home())
}

func (s *Server) handleWeapons(w http.ResponseWriter, r *http.Request) error {
	weapons, err := s.mongo.Weapons(r.Context())
	if err != nil {
		return err
	}

	return lib.WriteJSON(w, http.StatusOK, WeaponsResponse{Weapons: weapons})
}

func (s *Server) handleWeaponByName(w http.ResponseWriter, r *http.Request) error {
	name := chi.URLParam(r, "name")

	params, err := s.mongo.WeaponByName(r.Context(), name)
	if err != nil {
		return lib.InvalidRequest(name)
	}

	return lib.WriteJSON(w, http.StatusOK, params)
}

func (s *Server) handleWeaponsByCategory(w http.ResponseWriter, r *http.Request) error {
	category := chi.URLParam(r, "category")

	weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
	if err != nil {
		return lib.InvalidRequest(category)
	}

	return lib.Render(w, r, rearaspect.RearAspect(weapons))
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
		return lib.InvalidInsertData(req.Name)
	}

	return lib.WriteJSON(w, http.StatusOK, struct{}{})
}

func (s *Server) handleUpdateWeapon(w http.ResponseWriter, r *http.Request) error {
	name := chi.URLParam(r, "name")

	req := new(models.Params)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	if err := s.mongo.UpdateWeapon(r.Context(), name, req); err != nil {
		return lib.InvalidUpdateData(name)
	}

	return lib.WriteJSON(w, http.StatusOK, struct{}{})
}

func (s *Server) handleSearchWeapon(w http.ResponseWriter, r *http.Request) error {
	search := chi.URLParam(r, "search")

	weapons, err := s.mongo.SearchWeapon(r.Context(), search)
	if err != nil {
		return err
	}

	return lib.WriteJSON(w, http.StatusOK, WeaponsResponse{Weapons: weapons})
}
