package dto

import "time"

type ExpenseRequest struct {
	Amount     float64 `json:"amount" validate:"required,gt=0"`
	CategoryID uint    `json:"categoryID" validate:"required"`
	Note       string  `json:"note"`
}

type ExpenseResponse struct {
	ID         uint      `json:"id"`
	Amount     float64   `json:"amount"`
	UserID     uint      `json:"userID"`
	CategoryID uint      `json:"categoryID"`
	Note       string    `json:"note"`
	CreatedAt  time.Time `json:"createdAt"`
}

type ExpenseReportRequest struct {
	FromDate time.Time `json:"fromDate" validate:"required"`
	ToDate   time.Time `json:"toDate" validate:"required"`
}

type ExpenseReportResponse struct {
	CategoryID uint    `json:"categoryID"`
	Amount     float64 `json:"amount"`
}
