package repository

import (
	"expense_tracker/dto"
	"expense_tracker/model"
	"time"

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
	expense := r.db.Unscoped().Where("id = ? AND user_id = ?", id, userID).Delete(&model.Expense{})

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

func (r *ExpenseRepository) GetPaginatedExpenses(categoryID *uint, userID uint, offset, limit int) ([]model.Expense, int64, error) {
	var expenses []model.Expense
	var total int64

	query := r.db.Model(&model.Expense{}).Where("user_id = ?", userID)

	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&expenses).Error; err != nil {
		return nil, 0, err
	}

	return expenses, total, nil
}

func (r *ExpenseRepository) GetSummary(userID uint, fromDate, toDate time.Time) ([]dto.CategoryExpenseSummary, error) {
	var rows []dto.CategoryExpenseSummary
	if err := r.db.Table("categories").
		Select("categories.name AS category, COALESCE(SUM(expenses.amount),0) AS total_amount, COUNT(expenses.id) AS total_transactions").
		Joins("LEFT JOIN expenses ON categories.id = expenses.category_id AND expenses.created_at BETWEEN ? AND ?", fromDate, toDate).
		Where("categories.user_id = ?", userID).
		Group("categories.id").
		Order("total_amount DESC").
		Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *ExpenseRepository) GetYearlySummary(userID uint, year int) ([]dto.MonthlyCategoryExpense, error) {
	var rows []dto.MonthlyCategoryExpense
	if err := r.db.Raw(`
		WITH months AS (SELECT GENERATE_SERIES(1,12) AS month)
		SELECT m.month AS month, c.name AS category, COALESCE(SUM(e.amount),0) AS total_amount, COUNT(e.id) AS total_transactions
		FROM months m CROSS JOIN categories c 
		LEFT JOIN expenses e 
		ON c.id = e.category_id AND EXTRACT(MONTH FROM e.created_at) = m.month AND EXTRACT(YEAR FROM e.created_at) = ?
		WHERE c.user_id = ?
		GROUP BY m.month, c.id
		ORDER BY m.month ASC, c.id ASC

	`, year, userID).Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
