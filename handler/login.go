package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/harshitrajsinha/goserver-vanmango/models"
)

func GenerateToken(username string) (string, error) {

	expiration := time.Now().Add(24 * time.Hour)

	claims := &jwt.StandardClaims{
		ExpiresAt: expiration.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("some_value"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	valid := (credentials.Username == "admin" && credentials.Password == "harshit")

	if !valid {
		http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
		return
	}

	tokenString, err := GenerateToken(credentials.Username)

	if err != nil {
		http.Error(w, "Failed to generate token ", http.StatusInternalServerError)
		log.Println("Error generating token ", err)
		return
	}

	response := map[string]string{"token": tokenString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
