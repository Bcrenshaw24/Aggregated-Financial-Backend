package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if v != nil {
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("encode error: %v", err)
		}
	}
}
