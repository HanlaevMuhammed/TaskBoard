package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"type:varchar(255); not null"`
	Description string `gorm:"type:text"`
	Status      string `gorm:"type:varchar(20); default: 'todo'"`
}
