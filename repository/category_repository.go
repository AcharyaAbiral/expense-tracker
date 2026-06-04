package repository

import (
	"expense_tracker/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (r *CategoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) FindByID(id uint) (*model.Category, error) {
	var category model.Category

	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}
