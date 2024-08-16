package lib

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewApiError(status int, err error) APIError {
	return APIError{
		StatusCode: status,
		Msg:        err.Error(),
	}
}

func InvalidRequest(s string) APIError {
	return NewApiError(http.StatusBadRequest, fmt.Errorf("nothing found for %s", s))
}

func InvalidInsertData(s string) APIError {
	return NewApiError(http.StatusBadRequest, fmt.Errorf("%s already exists", s))
}

func InvalidUpdateData(s string) APIError {
	return NewApiError(http.StatusBadRequest, fmt.Errorf("%s doesn't exist", s))
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHTTP(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			if apiErr, ok := err.(APIError); ok {
				WriteJSON(w, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"statusCode": http.StatusInternalServerError,
					"msg":        "internal server error",
				}
				WriteJSON(w, http.StatusInternalServerError, errResp)
			}
			slog.Error("API error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
