package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null"`
	Email        string `gorm:"unique"`
	PasswordHash string `gorm:"not null" json:"-"`
}
