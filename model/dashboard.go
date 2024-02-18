package model

type DashboardShowResponse struct {
	Title    string `json:"title" validate:"required"`
	Total    int    `json:"total" validate:"required"`
	TotalStr string `json:"total_str" validate:"required"`
}
