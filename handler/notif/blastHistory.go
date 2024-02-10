package notifHandler

import (
	"class-reminder-be/config"
	"class-reminder-be/library"
	usecase "class-reminder-be/usecase/notif"
	"fmt"
	"net/http"
	"time"
)

func BlastHistory(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// do business logic
	data, err := usecase.BlastHistory()
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
		created_at := string(item["created_at"].([]uint8))
		id := string(item["id"].([]uint8))
		message := string(item["message"].([]uint8))
		user_success := string(item["user_success"].([]uint8))

		// Parse the input timestamp string
		layout := "2006-01-02 15:04:05"
		timestamp, err := time.Parse(layout, created_at)
		if err != nil {
			fmt.Println("Error parsing timestamp:", err)
			return
		}

		// Format the timestamp as per the desired format
		formattedTime := timestamp.Format("02/01/2006 15:04:05")

		// Create a map with string values
		data := map[string]interface{}{
			"created_at":     created_at,
			"id":             id,
			"message":        message,
			"user_success":   user_success,
			"created_at_str": formattedTime,
		}

		jsonData = append(jsonData, data)
	}
	_, responseJSON := library.SetResponse(config.RCSuccess, config.DescSuccess, jsonData)
	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
