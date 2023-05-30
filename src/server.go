package main

import (
	"github.com/labstack/echo/v4"
	"yandex-team.ru/bstask/config"
	db2 "yandex-team.ru/bstask/db"
	"yandex-team.ru/bstask/handler"
	"yandex-team.ru/bstask/routes"
)

func main() {
	e := setupServer()
	e.Logger.Fatal(e.Start(":8080"))
}

func setupServer() *echo.Echo {
	e := echo.New()
	cfg, err := config.NewConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}
	db, err := db2.InitDB(cfg)
	if err != nil {
		e.Logger.Fatal(e)
	}
	err = db2.Migrate(db)
	if err != nil {
		e.Logger.Fatal(err)
	}
	h := handler.NewHandler(db)
	routes.SetupRoutes(e, h)
	return e
}
