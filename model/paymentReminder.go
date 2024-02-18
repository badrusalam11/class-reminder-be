package model

type PaymentReminderResponse struct {
	Id              int    `json:"id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Nim             string `json:"nim" validate:"required"`
	TuitionFee      int    `json:"tuition_fee" validate:"required"`
	VaAccount       string `json:"va_account" validate:"required"`
	LastPaymentDate string `json:"last_payment_date" validate:"required"`
	DueDate         string `json:"due_date" validate:"required"`
}

type JobDetailResponse struct {
	IdEvent    int    `json:"id_event" validate:"required"`
	Title      string `json:"title" validate:"required"`
	IsJobExist bool   `json:"is_job_exist" validate:"required"`
	Schedule   string `json:"schedule" validate:"required"`
	JobEvery   string `json:"job_every" validate:"required"`
	JobName    string `json:"job_name" validate:"required"`
	JobId      string `json:"job_id" validate:"required"`
	ButtonText string `json:"button_text" validate:"required"`
}

type JobTriggerRequest struct {
	Nim     string `json:"nim" validate:"required"`
	IdEvent int    `json:"id_event" validate:"required"`
}
