package repository

import (
	"class-reminder-be/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, error) {
	// Define the claims for your token
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * config.JwtDuration).Unix(), // Token expiration time (1 hour from now)
	}

	// Create a new token with the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(config.JwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
