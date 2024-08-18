package repository

// Notification is a struct to represent the JSON data you expect to receive.
type FirebaseRequest struct {
	RegistrationIds []string             `json:"registration_ids"`
	Notification    FirebaseNotification `json:"notification"`
}

type WhatsappRequest struct {
	ApiKey  string `json:"api_key"`
	Sender  string `json:"sender"`
	Number  string `json:"number"`
	Message string `json:"message"`
}

type WhatsappNewRequest struct {
	AppKey  string `json:"appkey"`
	Authkey string `json:"authkey"`
	To      string `json:"to"`
	Message string `json:"message"`
	Sandbox string `json:"sandbox"`
}

type FirebaseNotification struct {
	Body     string `json:"body"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}
