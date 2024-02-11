package courseHandler

import (
	"class-reminder-be/config"
	"class-reminder-be/library"
	usecase "class-reminder-be/usecase/course"
	"net/http"
	"strconv"
)

func List(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// do business logic
	data, err := usecase.List()
	if err != nil {
		_, responseJSON := library.SetResponse(config.RCSnackbar, config.DescSnackbar, map[string]interface{}{})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
		return

	}

	var jsonData []map[string]interface{}
	for _, item := range data {
		// Convert byte arrays to strings
		title := string(item["title"].([]uint8))
		idString := string(item["id"].([]uint8))
		id, _ := strconv.Atoi(idString)
		// Create a map with string values
		data := map[string]interface{}{
			"title": title,
			"id":    id,
		}

		jsonData = append(jsonData, data)
	}
	_, responseJSON := library.SetResponse(config.RCSuccess, config.DescSuccess, jsonData)
	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
