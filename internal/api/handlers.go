package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/zeze322/wt-guided-weaponry/lib"
	"github.com/zeze322/wt-guided-weaponry/models"
	"github.com/zeze322/wt-guided-weaponry/views/agmsalh"
	"github.com/zeze322/wt-guided-weaponry/views/atgmautomatic"
	"github.com/zeze322/wt-guided-weaponry/views/atgmlosbr"
	"github.com/zeze322/wt-guided-weaponry/views/components/search"
	"github.com/zeze322/wt-guided-weaponry/views/home"
	"github.com/zeze322/wt-guided-weaponry/views/rearaspect"
)

const (
	irRearAspect   = "ir-rear-aspect"
	irAllAspect    = "ir-all-aspect"
	IrHeli         = "ir-heli"
	aamSarh        = "aam-sarh"
	aamArh         = "aamArh"
	aamMclosLosbr  = "aam-mclos-losbr"
	agmAutomatic   = "agm-automatic"
	agmSalh        = "agm-salh"
	agmSaclos      = "agm-saclos"
	agmMclos       = "agm-mclos"
	agmLosbr       = "agm-losbr"
	gbu            = "gbu"
	samIr          = "sam-ir"
	samSaclosLosbr = "sam-saclos-losbr"
	atgmMclos      = "atgm-mclos"
	atgmSaclos     = "atgm-saclos"
	atgmLosbr      = "atgm-losbr"
	atgmAutomatic  = "atgm-automatic"
	ashm           = "ashm"
)

type SearchResponse struct {
	Name []models.Name `json:"weapons"`
}

type WeaponsResponse struct {
	Weapons []*models.Params `json:"weapons"`
}

func (s *Server) handleWeapon(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "PUT":
		return s.handleUpdateWeapon(w, r)
	}

	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) error {

	return lib.Render(w, r, home.Home())
}

func (s *Server) handleCategories(w http.ResponseWriter, r *http.Request) error {
	categories, err := s.mongo.Categories(r.Context())
	if err != nil {
		return err
	}

	return lib.WriteJSON(w, http.StatusOK, categories)
}

func (s *Server) handleWeapons(w http.ResponseWriter, r *http.Request) error {
	if r.FormValue("search") != "" {
		return s.handleSearchWeapon(w, r)
	}

	weapons, err := s.mongo.Weapons(r.Context())
	if err != nil {
		return err
	}

	return lib.WriteJSON(w, http.StatusOK, WeaponsResponse{Weapons: weapons})
}

func (s *Server) handleWeaponsByCategory(w http.ResponseWriter, r *http.Request) error {
	category := r.FormValue("name")

	switch category {
	case irRearAspect:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, rearaspect.RearAspect(weapons))
	case irAllAspect:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, rearaspect.RearAspect(weapons))
	case agmSalh:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, agmsalh.AgmSalh(weapons))
	case atgmAutomatic:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, atgmautomatic.AtgmAutomatic(weapons))
	case atgmLosbr:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, atgmlosbr.AtgmLosbr(weapons))
	}

	return nil
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
	keyWord := r.FormValue("search")

	weapons, err := s.mongo.SearchWeapon(r.Context(), keyWord)
	if err != nil {
		return err
	}

	return lib.Render(w, r, search.SearchResult(weapons))
}
