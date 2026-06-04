package model

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	Amount float64

	UserID uint
	User   User `gorm:"constraint:OnDelete:CASCADE;"`

	CategoryID uint     //todo: non null constraint
	Category   Category `gorm:"constraint:OnDelete:CASCADE;"`

	Note string
}
