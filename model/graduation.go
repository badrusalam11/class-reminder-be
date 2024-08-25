package model

type GraduationShowResponse struct {
	Id            int64  `json:"id" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Nim           string `json:"nim" validate:"required"`
	Major         string `json:"major" validate:"required"`
	Is_registered string `json:"is_registered" validate:"required"`
}
