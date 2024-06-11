package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	middleware_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/middleware/interfaces"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"
	"github.com/gorilla/mux"
)

func JsonBodyDecoder(body io.ReadCloser, req any) error {
	if err := json.NewDecoder(body).Decode(&req); err != nil {
		log.Printf("Failed to decode request body: %v", err)
		return err
	}
	return nil
}

func SendResponse(w http.ResponseWriter, response interface{}, status int) {
	res, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to parse response: %v", err)
		http.Error(w, "Failed to parse response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)
}

func GetCurrentUser(r *http.Request) (*response.CurrentUser, error) {
	// Extract values from context
	userID, ok := r.Context().Value(middleware_intf.ContextKey.UserID).(int64)
	if !ok {
		return nil, errors.New("user id is required or invalid")
	}

	userRole, ok := r.Context().Value(middleware_intf.ContextKey.UserRole).(string)
	if !ok {
		return nil, errors.New("user role is required or invalid")
	}

	return &response.CurrentUser{
		ID:   userID,
		Role: userRole,
	}, nil
}

func UrlVars(r *http.Request, key string) string {
	return mux.Vars(r)[key]
}
