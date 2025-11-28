package database

import (
	"fmt"
	"log"
	"taskBoard_API/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	config.InitConfig()
	dbConf := config.Config.Database

	dsn := fmt.Sprintf(
		"host=%s, port=%d, user=%s, password=%s, dbname=%s, sslmode=%s",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name, dbConf.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: %v", err)
	}

	DB = db
	log.Println("База данных успешно подключена!")

}
