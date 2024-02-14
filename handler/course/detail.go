package courseHandler

import (
	"class-reminder-be/config"
	"class-reminder-be/library"
	"class-reminder-be/model"
	usecase "class-reminder-be/usecase/course"
	"encoding/json"
	"net/http"
)

func Detail(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body
	var request model.CourseDeleteRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// do business logic
	data, err := usecase.Detail(request)
	if err != nil {
		_, responseJSON := library.SetResponse(config.RCSnackbar, config.DescSnackbar, map[string]interface{}{})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
		return

	}

	_, responseJSON := library.SetResponse(config.RCSuccess, config.DescSuccess, data)
	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
