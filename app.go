package main

import "github.com/Daanooo/vams/internal/router"

type App struct {
	router *router.Router
}

func NewApp() *App {
	router := router.NewRouter()

	return &App{
		router: router,
	}
}

func (app App) Run() {
	app.router.Run()
}
