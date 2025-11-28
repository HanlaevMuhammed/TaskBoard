package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	User_id     uint      `gorm:"not null"`
	Title       string    `gorm:"type:varchar(255); not null"`
	Description string    `gorm:"type:text"`
	Status      string    `gorm:"type:varcher(20); default: 'todo'"`
	Created_at  time.Time `gorm:"autoCreateTime"`
	Updated_at  time.Time `gorm:"autoUpdateTime"`
}
