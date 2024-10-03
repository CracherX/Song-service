package main

import (
	app "github.com/CracherX/Song-service/pkg/song/app"
	"log"
)

// @title Реализация онлайн библиотеки песен
// @version 1.0
// @description Выполнялось как тестовое задание для Effective Mobile
// @host localhost:8080
// @BasePath /songs
func main() {
	App, err := app.New()
	if err != nil {
		log.Fatalf("Ошибка запуска приложения!")
	}
	App.Run()
}
