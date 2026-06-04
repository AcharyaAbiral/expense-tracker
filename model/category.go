package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)"`

	UserID uint
	User   User `gorm:"constraint:OnDelete:CASCADE;"`
}
