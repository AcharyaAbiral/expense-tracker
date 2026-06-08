package model

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Amount float64

	UserID uint `gorm:"index:idx_user_category,priority:1"`
	User   User `gorm:"constraint:OnDelete:CASCADE;"`

	CategoryID uint     `gorm:"index:idx_user_category,priority:2"`
	Category   Category `gorm:"constraint:OnDelete:CASCADE;"`

	Note string
}
