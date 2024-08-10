package lib

import (
	"encoding/json"
	"net/http"
)

type APIErr struct {
	Error string `json:"error"`
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHTTP(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIErr{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
