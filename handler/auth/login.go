package authHandler

import (
	"class-reminder-be/config"
	library "class-reminder-be/library"
	"class-reminder-be/model"
	usecase "class-reminder-be/usecase/auth"
	"encoding/json"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Decode the JSON body
	var request model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("before token")
	// do business logic
	token, errCode := usecase.Login(request)
	fmt.Println("token", token)
	fmt.Println("token", errCode)
	if errCode != "" {
		_, responseJSON := library.SetResponse(config.RCSnackbar, config.DescSnackbar, map[string]interface{}{})
		if errCode == "UZ" {
			_, responseJSON = library.SetResponse(config.RCSnackbar, "Username atau password salah", map[string]interface{}{})
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
		return

	}

	loginResponse := model.LoginResponse{
		Username: request.Username,
		Token:    token,
	}
	_, responseJSON := library.SetResponse(config.RCSuccess, config.DescSuccess, loginResponse)

	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
