package config

import (
	"log"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"sslmode"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type JWTConfig struct {
	Secret    string `yaml:"secret"`
	ExpiresIn int    `yaml:"expires_in"`
}

type AppConfig struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
}

var Config AppConfig

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./internal")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения файла config.yaml: %v", err)
	}

	Config.Database.Host = viper.GetString("database.host")
	Config.Database.Port = viper.GetInt("database.port")
	Config.Database.User = viper.GetString("database.user")
	Config.Database.Password = viper.GetString("database.password")
	Config.Database.Name = viper.GetString("database.name")
	Config.Database.SSLMode = viper.GetString("database.sslmode")

	Config.Server.Port = viper.GetInt("server.port")

	Config.JWT.Secret = viper.GetString("jwt.secret")
	Config.JWT.ExpiresIn = viper.GetInt("jwt.expires_in")

}
