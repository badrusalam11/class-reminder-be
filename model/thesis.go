package model

type ThesisShowResponse struct {
	Id                   int64  `json:"id" validate:"required"`
	Name                 string `json:"name" validate:"required"`
	Nim                  string `json:"nim" validate:"required"`
	Logbook              int    `json:"logbook" validate:"required"`
	Major                string `json:"major" validate:"required"`
	Last_attendance_date string `json:"last_attendance_date" validate:"required"`
	Is_eligable_grad     string `json:"Is_eligable_grad" validate:"required"`
	Is_reminder          bool   `json:"is_reminder" validate:"required"`
}
