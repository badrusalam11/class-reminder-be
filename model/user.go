package model

// Notification is a struct to represent the JSON data you expect to receive.
type RegisterUserRequest struct {
	Name       string `json:"name" validate:"required"`
	Nim        string `json:"nim" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	Major      string `json:"major" validate:"required"`
	TuitionFee int    `json:"tuition_fee" validate:"required"`
	// VaAccount       string `json:"va_account" validate:"required"`
	LastPaymentDate   string `json:"last_payment_date"`
	IsRegisGraduation int    `json:"is_regis_graduation" validate:"required"`
	Logbook           int    `json:"logbook" validate:"required"`
}

type EditUserRequest struct {
	Name              string `json:"name" validate:"required"`
	Nim               string `json:"nim" validate:"required"`
	Phone             string `json:"phone" validate:"required"`
	Major             string `json:"major" validate:"required"`
	TuitionFee        int    `json:"tuition_fee"`
	VaAccount         string `json:"va_account"`
	LastPaymentDate   string `json:"last_payment_date"`
	IsRegisGraduation int    `json:"is_regis_graduation" validate:"required"`
	Logbook           int    `json:"logbook" validate:"required"`
}

type DetailUserRequest struct {
	Nim string `json:"nim" validate:"required"`
}

type DeleteUserRequest struct {
	Nim string `json:"nim" validate:"required"`
}

type ClassDetail struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Result struct {
	Major             string `json:"major"`
	Name              string `json:"name"`
	NIM               string `json:"nim"`
	NoHP              string `json:"no_hp"`
	TuitionFee        int    `json:"tuition_fee"`
	VaAccount         string `json:"va_account"`
	LastPaymentDate   string `json:"last_payment_date"`
	IsRegisGraduation int    `json:"is_regis_graduation" validate:"required"`
	IsDoneThesis      int    `json:"is_done_thesis" validate:"required"`
	Logbook           int    `json:"logbook" validate:"required"`
}
