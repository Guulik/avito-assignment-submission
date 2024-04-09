package app

import (
	"Avito_trainee_assignment/internal/api"
	"Avito_trainee_assignment/internal/config"
	service "Avito_trainee_assignment/internal/service/bannerSvc"
	"Avito_trainee_assignment/internal/storage/postgresql"
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
)

type App struct {
	api     *api.Api
	svc     *service.Service
	storage *postgresql.Storage
	echo    *echo.Echo
}

func New(log *slog.Logger, cfg *config.Config) *App {
	app := &App{}

	app.echo = echo.New()

	db, err := postgresql.InitPostgres(cfg)
	if err != nil {
		log.Error("failed to connect to PostgresSQL", err)
	}
	err = postgresql.CreateTable(db)
	if err != nil {
		log.Error("failed to connect to create table in DB", err)
	}

	app.storage = postgresql.New(log, db)

	app.svc = service.New(log, app.storage)

	app.api = api.New(log, app.svc)

	app.echo.GET("/user_banner/get", app.api.GetUserBanner)
	app.echo.GET("/banner", app.api.GetBanner)
	app.echo.POST("/banner", app.api.CreateBanner)
	app.echo.PATCH("/banner/:id", app.api.PatchBanner)
	app.echo.DELETE("/banner/:id", app.api.DeleteBanner)

	return app
}

func (a *App) Run() error {
	fmt.Println("server running")

	err := a.echo.Start(":4444")
	if err != nil {
		return err
	}

	return nil
}
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}
