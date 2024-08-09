package lib

import (
	"encoding/json"
	"net/http"
)

type ApiErr struct {
	Error string `json:"error"`
}

type ApiFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHTTP(fn ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiErr{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
