package app

import (
	"fmt"
	"github.com/CracherX/Song-service/internal/song/client"
	"github.com/CracherX/Song-service/internal/song/config"
	"github.com/CracherX/Song-service/internal/song/endpoints"
	"github.com/CracherX/Song-service/internal/song/logger"
	"github.com/CracherX/Song-service/internal/song/middleware"
	"github.com/CracherX/Song-service/internal/song/router"
	"github.com/CracherX/Song-service/internal/song/services"
	"github.com/CracherX/Song-service/internal/song/storage/db"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger
	DB     *gorm.DB
	Router *mux.Router
}

func New() (*App, error) {
	var app App
	var err error
	app.Config = config.MustLoad()
	app.Logger = logger.MustInit(app.Config.Server.Debug)
	app.Logger.Debug(
		"Подключение к базе данных...",
		zap.String("Имя", app.Config.Database.Name),
		zap.String("Хост", app.Config.Database.Host),
		zap.String("Пользователь", app.Config.Database.User),
	)
	app.DB, err = db.Connect(app.Config)
	if err != nil {
		app.Logger.Error("Ошибка подключения к базе данных: ", zap.Error(err))
		return nil, err
	}
	app.Logger.Debug("Успешное подключение к БД")

	ss := services.NewSongsService(app.DB)
	sc := client.NewClient("http://127.0.0.1:8080")
	ep := endpoints.NewSongsEndpoint(ss, sc, app.Logger)

	app.Router = router.Setup()

	app.Router.Use(middleware.Validate(validator.New()))

	router.Songs(app.Router, ep)

	return &app, nil
}

// Run запуск приложения.
func (a *App) Run() {
	a.Logger.Info("Запуск приложения", zap.String("Приложение", a.Config.Server.AppName))
	a.Logger.Debug("Запущен режим отладки для терминала!")
	err := http.ListenAndServe(a.Config.Server.Port, a.Router)
	if err != nil {
		fmt.Println(err)
		a.Logger.Error("Ошибка запуска сервера")
	}
}
