package model

// Notification is a struct to represent the JSON data you expect to receive.
type NotifRequest struct {
	EventId int `json:"event_id" validate:"required"`
}

// Response is a struct to represent the response data.
type NotifResponse struct {
	Status string `json:"status"`
	// Message string `json:"message"`
}
