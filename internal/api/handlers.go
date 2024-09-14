package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/zeze322/wt-guided-weaponry/lib"
	"github.com/zeze322/wt-guided-weaponry/models"
	"github.com/zeze322/wt-guided-weaponry/views/aamarh"
	"github.com/zeze322/wt-guided-weaponry/views/aammcloslosbr"
	"github.com/zeze322/wt-guided-weaponry/views/aamsarh"
	"github.com/zeze322/wt-guided-weaponry/views/agmautomatic"
	"github.com/zeze322/wt-guided-weaponry/views/agmlosbr"
	"github.com/zeze322/wt-guided-weaponry/views/agmmclos"
	"github.com/zeze322/wt-guided-weaponry/views/agmsaclos"
	"github.com/zeze322/wt-guided-weaponry/views/agmsalh"
	"github.com/zeze322/wt-guided-weaponry/views/allaspect"
	"github.com/zeze322/wt-guided-weaponry/views/ashm"
	"github.com/zeze322/wt-guided-weaponry/views/atgmautomatic"
	"github.com/zeze322/wt-guided-weaponry/views/atgmlosbr"
	"github.com/zeze322/wt-guided-weaponry/views/atgmmclos"
	"github.com/zeze322/wt-guided-weaponry/views/atgmsaclos"
	"github.com/zeze322/wt-guided-weaponry/views/components/search"
	"github.com/zeze322/wt-guided-weaponry/views/gbu"
	"github.com/zeze322/wt-guided-weaponry/views/home"
	"github.com/zeze322/wt-guided-weaponry/views/irheli"
	"github.com/zeze322/wt-guided-weaponry/views/rearaspect"
	"github.com/zeze322/wt-guided-weaponry/views/samir"
	"github.com/zeze322/wt-guided-weaponry/views/samsacloslosbr"
)

const (
	irRearAspect   = "ir-rear-aspect"
	irAllAspect    = "ir-all-aspect"
	irHeli         = "ir-heli"
	aamSarh        = "aam-sarh"
	aamArh         = "aam-arh"
	aamMclosLosbr  = "aam-mclos-losbr"
	agmAutomatic   = "agm-automatic"
	agmSalh        = "agm-salh"
	agmSaclos      = "agm-saclos"
	agmMclos       = "agm-mclos"
	agmLosbr       = "agm-losbr"
	gbuc           = "gbu"
	samIr          = "sam-ir"
	samSaclosLosbr = "sam-saclos-losbr"
	atgmMclos      = "atgm-mclos"
	atgmSaclos     = "atgm-saclos"
	atgmLosbr      = "atgm-losbr"
	atgmAutomatic  = "atgm-automatic"
	ashms          = "ashm"
)

type SearchResponse struct {
	Name []models.Name `json:"weapons"`
}

type WeaponsResponse struct {
	Weapons []*models.Params `json:"weapons"`
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
		return lib.Render(w, r, allaspect.AllAspect(weapons))
	case irHeli:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, irheli.IrHeli(weapons))
	case aamSarh:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, aamsarh.AamSarh(weapons))
	case aamArh:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, aamarh.AamArh(weapons))
	case aamMclosLosbr:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, aammcloslosbr.AamMclosLosbr(weapons))
	case agmAutomatic:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, agmautomatic.AgmAutomatic(weapons))
	case agmSalh:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, agmsalh.AgmSalh(weapons))
	case agmSaclos:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, agmsaclos.AgmSaclos(weapons))
	case agmMclos:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, agmmclos.AgmMclos(weapons))
	case agmLosbr:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, agmlosbr.AgmLosbr(weapons))
	case gbuc:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, gbu.Gbu(weapons))
	case samIr:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, samir.SamIr(weapons))
	case samSaclosLosbr:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, samsacloslosbr.SamSaclosLosbr(weapons))
	case atgmMclos:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, atgmmclos.AtgmMclos(weapons))
	case atgmSaclos:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, atgmsaclos.AtgmSaclos(weapons))
	case atgmLosbr:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, atgmlosbr.AtgmLosbr(weapons))
	case atgmAutomatic:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, atgmautomatic.AtgmAutomatic(weapons))
	case ashms:
		weapons, err := s.mongo.WeaponsByCategory(r.Context(), category)
		if err != nil {
			return err
		}
		return lib.Render(w, r, ashm.Ashm(weapons))
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
	if r.Method != "PUT" {
		return fmt.Errorf("method not allowed: %s", r.Method)
	}

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

	if keyWord == "" {
		return nil
	}

	weapons, err := s.mongo.SearchWeapon(r.Context(), keyWord)
	if err != nil {
		return err
	}

	return lib.Render(w, r, search.SearchResult(weapons))
}
