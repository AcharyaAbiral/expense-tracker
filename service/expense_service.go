package service

import (
	"expense_tracker/dto"
	"expense_tracker/mapper"
	"expense_tracker/model"
	"expense_tracker/repository"
	"fmt"
	"time"
)

type ExpenseService struct {
	repo            *repository.ExpenseRepository
	categoryService *CategoryService
}

func NewExpenseService(repo *repository.ExpenseRepository, categoryService *CategoryService) *ExpenseService {
	return &ExpenseService{repo, categoryService}
}

func (s *ExpenseService) Create(reqDto *dto.ExpenseRequest, userID uint) (*model.Expense, error) {

	expense := mapper.ToExpense(reqDto)
	expense.UserID = userID

	if _, err := s.categoryService.FindByID(expense.CategoryID); err != nil {
		return nil, err
	}

	if err := s.repo.Create(expense); err != nil {
		return nil, err
	}

	return expense, nil
}

func (s *ExpenseService) Update(reqDto *dto.ExpenseRequest, id uint, userID uint) (*model.Expense, error) {
	expense, err := s.FindByIDAndUserID(id, userID)
	if err != nil {
		return nil, err
	}

	if _, err := s.categoryService.FindByID(expense.CategoryID); err != nil {
		return nil, err
	}

	expense.Amount = reqDto.Amount
	expense.CategoryID = reqDto.CategoryID
	expense.Note = reqDto.Note

	if err := s.repo.Update(expense); err != nil {
		return nil, err
	}

	return expense, nil
}

func (s *ExpenseService) DeleteByIDAndUserID(id, userID uint) error {
	return s.repo.DeleteByIDAndUserID(id, userID)
}

func (s *ExpenseService) FindByIDAndUserID(id uint, userID uint) (*model.Expense, error) {
	expense, err := s.repo.FindByIDAndUserID(id, userID)
	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (s *ExpenseService) GetExpenses(categoryID *uint, userID uint, paginationInput dto.PaginationInput) (*dto.PaginatedResponse[dto.ExpenseResponse], error) {
	offset := (paginationInput.Page - 1) * paginationInput.Limit

	expenses, total, err := s.repo.GetPaginatedExpenses(categoryID, userID, offset, paginationInput.Limit)

	if err != nil {
		return nil, err
	}

	expensesResponse := make([]dto.ExpenseResponse, 0, len(expenses))

	for _, e := range expenses {
		expensesResponse = append(expensesResponse, *mapper.ToExpenseResponse(&e))
	}

	return &dto.PaginatedResponse[dto.ExpenseResponse]{
		Data:  expensesResponse,
		Page:  paginationInput.Page,
		Limit: paginationInput.Limit,
		Total: total,
	}, nil
}

func (s *ExpenseService) GetSummary(userID uint, fromDate, toDate time.Time) (*dto.ExpenseSummaryResponse, error) {
	if toDate.Before(fromDate) {
		return nil, fmt.Errorf("to must be after from")
	}

	categorySummaries, err := s.repo.GetSummary(userID, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	totalTransactions := 0
	totalAmount := 0.0
	for _, s := range categorySummaries {
		totalTransactions += s.TotalTransactions
		totalAmount += s.TotalAmount
	}

	return &dto.ExpenseSummaryResponse{
		TotalAmount:       totalAmount,
		TotalTransactions: totalTransactions,
		ExpenseSummaries:  categorySummaries,
	}, nil
}

func (s *ExpenseService) GetYearlySummary(userID uint, year int) (*dto.YearlyExpenseSummaryResponse, error) {
	currentYear := time.Now().Year()
	if year < 2020 || year > currentYear {
		return nil, fmt.Errorf("invalid year")
	}

	expenses, err := s.repo.GetYearlySummary(userID, year)
	if err != nil {
		return nil, err
	}

	monthlySummaries := make([]dto.MonthlyCategoryExpenseSummaryResponse, 12)
	var yearlySummary dto.YearlyExpenseSummaryResponse

	for _, e := range expenses {
		idx := e.Month - 1
		monthlySummaries[idx].Month = e.Month
		monthlySummaries[idx].TotalAmount += e.TotalAmount
		monthlySummaries[idx].TotalTransactions += e.TotalTransactions
		monthlySummaries[idx].ExpenseSummaries = append(monthlySummaries[idx].ExpenseSummaries, dto.CategoryExpenseSummary{
			Category:          e.Category,
			TotalAmount:       e.TotalAmount,
			TotalTransactions: e.TotalTransactions,
		})
		yearlySummary.TotalAmount += e.TotalAmount
		yearlySummary.TotalTransactions += e.TotalTransactions

	}

	yearlySummary.ExpenseSummaries = monthlySummaries
	return &yearlySummary, nil

}
