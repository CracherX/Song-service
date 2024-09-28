package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// Config структура конфигурации приложения
type Config struct {
	Server struct {
		AppName string
		Port    string
		Debug   bool
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SslMode  string
	}
}

// MustLoad загружает конфигурацию приложения
func MustLoad() (config *Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации .env")
	}

	// Заполнение структуры Server
	config.Server.AppName = os.Getenv("APP_NAME")
	config.Server.Port = os.Getenv("SERVER_PORT")

	// Преобразование строки из .env в boolean
	dbg, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatalf("Значение DEBUG должно принимать True/False")
	}
	config.Server.Debug = dbg

	// Заполнение структуры Database
	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Database.User = os.Getenv("DB_USER")
	config.Database.Password = os.Getenv("DB_PASSWORD")
	config.Database.Name = os.Getenv("DB_NAME")
	config.Database.SslMode = os.Getenv("DB_SSLMODE")

	return config
}
