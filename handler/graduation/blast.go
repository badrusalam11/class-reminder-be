package graduationHandler

import (
	"class-reminder-be/config"
	"class-reminder-be/library"
	usecase "class-reminder-be/usecase/graduation"
	"fmt"
	"net/http"
)

func Blast(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// do business logic
	err := usecase.Blast()
	// empty data
	data := map[string]interface{}{}
	if err != nil {
		fmt.Println(err)
		_, responseJSON := library.SetResponse(config.RCSnackbar, config.DescSnackbar, data)
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
