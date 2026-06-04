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

type CategoryExpenseSummary struct {
	Category          string  `json:"category" gorm:"column:category"`
	TotalAmount       float64 `json:"totalAmount" gorm:"column:total_amount"`
	TotalTransactions int     `json:"totalTransactions" gorm:"column:total_transactions"`
}

type MonthlyCategoryExpense struct {
	Month int `json:"month" gorm:"column:month"`
	CategoryExpenseSummary
}

type MonthlyCategoryExpenseSummaryResponse struct {
	Month             int                      `json:"month"`
	TotalAmount       float64                  `json:"totalAmount"`
	TotalTransactions int                      `json:"totalTransactions"`
	ExpenseSummaries  []CategoryExpenseSummary `json:"expenseSummaries"`
}

type ExpenseSummaryResponse struct {
	TotalAmount       float64                  `json:"totalAmount"`
	TotalTransactions int                      `json:"totalTransactions"`
	ExpenseSummaries  []CategoryExpenseSummary `json:"expenseSummaries"`
}

type YearlyExpenseSummaryResponse struct {
	TotalAmount       float64                                 `json:"totalAmount"`
	TotalTransactions int                                     `json:"totalTransactions"`
	ExpenseSummaries  []MonthlyCategoryExpenseSummaryResponse `json:"expenseSummaries"`
}
