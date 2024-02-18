package notifHandler

import (
	"class-reminder-be/config"
	"class-reminder-be/library"
	"class-reminder-be/model"
	usecase "class-reminder-be/usecase/notif"
	"encoding/json"
	"net/http"
)

func Send(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body
	var notification model.NotifRequest
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// You can process the received data here and generate a response
	// For example, you can create a simple response message.
	// response := model.NotifResponse{
	// 	Status: "Success",
	// 	// Message: "Received message: " + notification.Message,
	// }

	// do business logic
	send, err := usecase.Send(notification.EventId)
	if err != nil || send == "" {
		_, responseJSON := library.SetResponse(config.RCSnackbar, config.DescSnackbar, map[string]interface{}{})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
		return

	}

	_, responseJSON := library.SetResponse(config.RCSuccess, config.DescSuccess, map[string]interface{}{})

	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
	// if send == "" {
	// 	fmt.Println(err)
	// 	response = model.NotifResponse{
	// 		Status: "Failed",
	// 		// Message: "Received message: " + notification.Message,
	// 	}

	// }

	// // Convert the response struct to JSON
	// responseJSON, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// // Set the Content-Type header and write the response
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(responseJSON)
}
