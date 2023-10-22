package model

type Response struct {
	Code        string      `json:"code" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Data        interface{} `json:"data" validate:"required"`
}
