package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"omegaui.io/edgeflow/entities"
)

var secretKey, _ = os.LookupEnv("JWT_SECRET")

var InvalidJWT = errors.New("invalid jwt")

func generateAccessToken(u *entities.UserEntity) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   u.ID,
		"username": u.Name,
		"exp":      time.Now().Add(time.Minute * 10).Unix(),
	})
	return token.SignedString(secretKey)
}

func verifyAccessToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return InvalidJWT
	}
	return nil
}

func Listen() {
	mux := http.NewServeMux()
	mux.HandleFunc("/generate-access-token", func(w http.ResponseWriter, r *http.Request) {
		var user entities.UserEntity
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		token, err := generateAccessToken(&user)
		if err != nil {
			http.Error(w, "Cannot create access token", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(entities.TokenEntity{
			Token: token,
		})
	})
	mux.HandleFunc("/verify-access-token", func(w http.ResponseWriter, r *http.Request) {
		var token entities.TokenEntity
		err := json.NewDecoder(r.Body).Decode(&token)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		err = verifyAccessToken(token.Token)
		status := true
		if err != nil {
			status = false
		}
		json.NewEncoder(w).Encode(struct {
			Success bool `json:"success"`
		}{
			Success: status,
		})
	})
	http.ListenAndServe(":3000", mux)
}
