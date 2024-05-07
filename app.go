package main

import (
	"github.com/Daanooo/vams/internal/router"
	"github.com/joho/godotenv"
)

type App struct {
	router *router.Router
}

func NewApp() (*App, error) {
	router := router.NewRouter()

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &App{
		router: router,
	}, nil
}

func (app App) Run() {
	app.router.Run()
}
