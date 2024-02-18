package model

type EventCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Schedule    string `json:"schedule" validate:"required"`
	IdEventType int    `json:"id_event_type" validate:"required"`
	JobEvery    string `json:"job_every"`
}

type EventEditRequest struct {
	Id          int    `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Schedule    string `json:"schedule" validate:"required"`
	IdEventType int    `json:"id_event_type" validate:"required"`
	JobEvery    string `json:"job_every"`
}

type EventCreateResponse struct {
	Status string `json:"status"`
}
