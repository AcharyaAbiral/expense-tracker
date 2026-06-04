package model

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	Amount     float64
	UserID     uint
	CategoryID uint
	Note       string
}
