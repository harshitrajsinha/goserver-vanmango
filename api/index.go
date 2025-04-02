package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	rootMessage := struct {
		Code    int
		Message string
	}{
		Code:    http.StatusOK,
		Message: "Server is functioning",
	}

	json.NewEncoder(w).Encode(rootMessage)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create router using gorilla/mux
	var router *mux.Router = mux.NewRouter()

	// Create routes
	router.HandleFunc("/", handleHomeRoute).Methods("GET")

	// create server
	fmt.Println("Listening at PORT ", port)
	router.ServeHTTP(w, r)
}
