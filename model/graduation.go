package model

type GraduationShowResponse struct {
	Id                   int64  `json:"id" validate:"required"`
	Name                 string `json:"name" validate:"required"`
	Nim                  string `json:"nim" validate:"required"`
	Major                string `json:"major" validate:"required"`
	Is_registered        bool   `json:"is_registered" validate:"required"`
	Is_registered_string string `json:"is_registered_string" validate:"required"`
}

type GraduationSendRequest struct {
	Nim string `json:"nim" validate:"required"`
}
