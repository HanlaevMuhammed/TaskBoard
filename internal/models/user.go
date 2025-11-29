package models

type User struct {
	Email        string `gorm:"uniqueIndex; not null"`
	PasswordHash string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Tasks        []Task
}
