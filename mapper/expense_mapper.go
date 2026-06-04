package mapper

import (
	"expense_tracker/dto"
	"expense_tracker/model"
)

func ToExpense(reqDto *dto.ExpenseRequest) *model.Expense {
	return &model.Expense{
		Amount:     reqDto.Amount,
		CategoryID: reqDto.CategoryID,
		Note:       reqDto.Note,
	}
}

func ToExpenseResponse(expense *model.Expense) *dto.ExpenseResponse {
	return &dto.ExpenseResponse{
		ID:         expense.ID,
		Amount:     expense.Amount,
		UserID:     expense.UserID,
		CategoryID: expense.CategoryID,
		Note:       expense.Note,
		CreatedAt:  expense.CreatedAt,
	}
}
