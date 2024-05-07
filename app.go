package main

import (
	"github.com/Daanooo/vams/internal/database"
	"github.com/Daanooo/vams/internal/router"
	"github.com/joho/godotenv"
)

type App struct {
	router *router.Router
}

func NewApp() (*App, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	db, err := database.NewMysql()
	if err != nil {
		return nil, err
	}

	router := router.NewRouter(db)

	return &App{
		router: router,
	}, nil
}

func (app App) Run() {
	app.router.Run()
}
