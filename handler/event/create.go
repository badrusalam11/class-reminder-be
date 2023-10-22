package eventHandler

import (
	"class-reminder-be/model"
	usecase "class-reminder-be/usecase/event"
	"encoding/json"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body
	var request model.EventCreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// You can process the received data here and generate a response
	// For example, you can create a simple response message.
	response := model.EventCreateResponse{
		Status: "Success",
		// Message: "Received message: " + notification.Message,
	}

	// do business logic
	send, err := usecase.Create(request)
	if send == "" {
		response = model.EventCreateResponse{
			Status: "Failed",
			// Message: "Received message: " + notification.Message,
		}
	}

	// Convert the response struct to JSON
	responseJSON, err := json.Marshal(&response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
