package graduationHandler

import (
	"class-reminder-be/config"
	"class-reminder-be/library"
	usecase "class-reminder-be/usecase/graduation"
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// do business logic
	data, err := usecase.Show()
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
