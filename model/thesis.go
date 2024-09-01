package model

type ThesisShowResponse struct {
	Id                   int64  `json:"id" validate:"required"`
	Name                 string `json:"name" validate:"required"`
	Nim                  string `json:"nim" validate:"required"`
	Supervisor           string `json:"supervisor" validate:"required"`
	Last_attendance_date string `json:"last_attendance_date" validate:"required"`
	Is_attend_this_week  bool   `json:"is_attend_this_week" validate:"required"`
}
