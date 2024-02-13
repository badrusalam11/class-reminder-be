package model

type CourseCrateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Schedule    string `json:"schedule" validate:"required"`
	Day         string `json:"day" validate:"required"`
}
