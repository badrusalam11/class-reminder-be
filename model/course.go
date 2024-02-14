package model

type CourseCrateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Schedule    string `json:"schedule" validate:"required"`
	Day         string `json:"day" validate:"required"`
}
type CourseEditRequest struct {
	Id          int64  `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Schedule    string `json:"schedule" validate:"required"`
	Day         string `json:"day" validate:"required"`
}
type CourseDeleteRequest struct {
	Id int64 `json:"id" validate:"required"`
}

type CourseShowResponse struct {
	Id              int64  `json:"id" validate:"required"`
	Title           string `json:"title" validate:"required"`
	Description     string `json:"description" validate:"required"`
	Schedule        string `json:"schedule" validate:"required"`
	CourseDay       string `json:"course_day" validate:"required"`
	CourseDayPrefix string `json:"course_day_prefix" validate:"required"`
	ReminderDay     string `json:"reminder_day" validate:"required"`
}
