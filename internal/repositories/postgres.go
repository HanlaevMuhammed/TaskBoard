package storage

import (
	"taskBoard_API/database"

	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return database.DB
}
