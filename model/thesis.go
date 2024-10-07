package model

type ThesisShowResponse struct {
	Id                   int64  `json:"id" validate:"required"`
	Name                 string `json:"name" validate:"required"`
	Nim                  string `json:"nim" validate:"required"`
	Logbook              int    `json:"logbook" validate:"required"`
	Major                string `json:"major" validate:"required"`
	Last_attendance_date string `json:"last_attendance_date" validate:"required"`
	// Is_attend_this_week  bool   `json:"is_attend_this_week" validate:"required"`
	Is_regis_graduation        bool   `json:"is_regis_graduation" validate:"required"`
	Is_regis_graduation_string string `json:"is_regis_graduation_string" validate:"required"`
	Is_reminder                bool   `json:"is_reminder" validate:"required"`
}
