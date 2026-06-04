package repository

import (
	"expense_tracker/model"

	"gorm.io/gorm"
)

type ExpenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{db}
}

func (r *ExpenseRepository) Create(expense *model.Expense) error {
	return r.db.Create(expense).Error
}

func (r *ExpenseRepository) Update(expense *model.Expense) error {
	return r.db.Save(expense).Error
}

func (r *ExpenseRepository) DeleteByIDAndUserID(id, userID uint) error {
	expense := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Expense{})

	if expense.Error != nil {
		return expense.Error
	}

	if expense.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *ExpenseRepository) FindByIDAndUserID(id uint, userID uint) (*model.Expense, error) {
	var expense model.Expense

	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&expense).Error; err != nil {
		return nil, err
	}

	return &expense, nil
}

func (r *ExpenseRepository) GetPaginatedExpenses(userID uint, offset, limit int) ([]model.Expense, int64, error) {
	var expenses []model.Expense
	var total int64

	if err := r.db.Model(&model.Expense{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&expenses).Error; err != nil {
		return nil, 0, err
	}

	return expenses, total, nil
}
