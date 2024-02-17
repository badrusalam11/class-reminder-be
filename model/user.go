package model

// Notification is a struct to represent the JSON data you expect to receive.
type RegisterUserRequest struct {
	Name       string `json:"name" validate:"required"`
	Nim        string `json:"nim" validate:"required"`
	Class      []int  `json:"class" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	Major      string `json:"major" validate:"required"`
	TuitionFee int    `json:"tuition_fee" validate:"required"`
	VaAccount  string `json:"va_account" validate:"required"`
}

type EditUserRequest struct {
	Name       string `json:"name" validate:"required"`
	Nim        string `json:"nim" validate:"required"`
	Class      []int  `json:"class" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	Major      string `json:"major" validate:"required"`
	TuitionFee int    `json:"tuition_fee"`
	VaAccount  string `json:"va_account"`
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
	Major      string        `json:"major"`
	Name       string        `json:"name"`
	NIM        string        `json:"nim"`
	NoHP       string        `json:"no_hp"`
	Class      []ClassDetail `json:"class"`
	ClassArr   []int         `json:"class_arr"`
	TuitionFee int           `json:"tuition_fee"`
	VaAccount  string        `json:"va_account"`
}
