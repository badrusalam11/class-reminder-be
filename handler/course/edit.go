package courseHandler

import (
	"class-reminder-be/config"
	"class-reminder-be/library"
	"class-reminder-be/model"
	usecase "class-reminder-be/usecase/course"
	"encoding/json"
	"net/http"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body
	var request model.CourseEditRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// do business logic
	err = usecase.Edit(request)
	if err != nil {
		_, responseJSON := library.SetResponse(config.RCSnackbar, config.DescSnackbar, map[string]interface{}{})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
		return

	}

	// var jsonData []map[string]interface{}
	// for _, item := range data {
	// 	// Convert byte arrays to strings
	// 	title := string(item["title"].([]uint8))
	// 	idString := string(item["id"].([]uint8))
	// 	id, _ := strconv.Atoi(idString)
	// 	// Create a map with string values
	// 	data := map[string]interface{}{
	// 		"title": title,
	// 		"id":    id,
	// 	}

	// 	jsonData = append(jsonData, data)
	// }
	_, responseJSON := library.SetResponse(config.RCSuccess, config.DescSuccess, map[string]interface{}{})
	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
