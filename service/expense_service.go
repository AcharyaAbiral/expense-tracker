package service

import (
	"expense_tracker/dto"
	"expense_tracker/mapper"
	"expense_tracker/model"
	"expense_tracker/repository"
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

func (s *ExpenseService) GetExpenses(userID uint, paginationInput dto.PaginationInput) (*dto.PaginatedResponse[dto.ExpenseResponse], error) {
	offset := (paginationInput.Page - 1) * paginationInput.Limit

	expenses, total, err := s.repo.GetPaginatedExpenses(userID, offset, paginationInput.Limit)

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
