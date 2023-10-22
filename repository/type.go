package repository

// Notification is a struct to represent the JSON data you expect to receive.
type FirebaseRequest struct {
	RegistrationIds []string             `json:"registration_ids"`
	Notification    FirebaseNotification `json:"notification"`
}

type FirebaseNotification struct {
	Body     string `json:"body"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}
