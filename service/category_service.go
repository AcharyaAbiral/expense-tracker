package service

import (
	"expense_tracker/dto"
	"expense_tracker/mapper"
	"expense_tracker/model"
	"expense_tracker/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo}
}

func (s *CategoryService) Create(reqDto *dto.CategoryRequest, userID uint) (*model.Category, error) {
	category := mapper.ToCategory(reqDto)
	category.UserID = userID

	if err := s.repo.Create(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) FindByID(id uint) (*model.Category, error) {
	return s.repo.FindByID(id)
}
