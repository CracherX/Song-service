package app

import "github.com/CracherX/Song-service/internal/song/config"

type App struct {
	Config *config.Config
}

func New() (app *App) {
	app.Config = config.MustLoad()

	return app
}
