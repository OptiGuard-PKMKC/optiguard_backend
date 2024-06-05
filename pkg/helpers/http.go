package helpers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
