package test

import (
	"encoding/json"
	"net/http"
)

func HandleHomeRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Server is functioning"})
}
