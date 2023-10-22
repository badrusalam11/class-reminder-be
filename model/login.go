package model

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Username string `json:"username" validate:"required"`
	Token    string `json:"token" validate:"required"`
}
