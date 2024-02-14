package middleware

import (
	// Other imports...
	"class-reminder-be/config"
	"class-reminder-be/library"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the request's headers, typically from the "Authorization" header.
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Your JWT secret key. Replace this with your actual secret key.
		secretKey := []byte(config.JwtKey)

		// Parse and validate the JWT token.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method and return the secret key.
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return secretKey, nil
		})

		if err != nil {
			_, responseJSON := library.SetResponse(config.RCSession, config.DescSession, []map[string]interface{}{})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseJSON)
			fmt.Println("token jwt", token)
			fmt.Println("token string", tokenString)
			fmt.Println(err)
			// http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid.
		if token.Valid {
			// If the token is valid, call the next handler in the chain.
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
	})
}
