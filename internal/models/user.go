package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex; not null"`
	PasswordHash string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Tasks        []Task `gorm:"foreignKey:UserID"`
}
