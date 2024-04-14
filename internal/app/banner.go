package app

import (
	"Avito_trainee_assignment/config"
	"Avito_trainee_assignment/internal/api"
	service "Avito_trainee_assignment/internal/service/bannerSvc"
	"Avito_trainee_assignment/internal/storage/postgresql"
	"Avito_trainee_assignment/internal/storage/redis"
	"context"
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"
)

type App struct {
	api     *api.Api
	svc     *service.Service
	storage *postgresql.Storage
	echo    *echo.Echo
	cache   *redis.Cache
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

	rds := redis.InitRedis(cfg)

	app.storage = postgresql.New(log, db)

	app.cache = redis.New(log, rds, cfg)

	app.svc = service.New(log, app.storage, app.cache)

	app.api = api.New(log, app.svc)

	app.echo.GET("/user_banner", app.api.GetUserBanner)
	app.echo.GET("/banner", app.api.GetBanner)
	app.echo.POST("/banner", app.api.CreateBanner)
	app.echo.PATCH("/banner/:id", app.api.PatchBanner)
	app.echo.DELETE("/banner/:id", app.api.DeleteBanner)
	app.echo.DELETE("/banner", app.api.DeleteBanner)

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

func (a *App) Stop() error {
	fmt.Println("stopping server..." + " op = app.Stop")
	ctx := context.Background()
	if err := a.echo.Shutdown(ctx); err != nil {
		fmt.Println("failed to gracefully stop server")
		return err
	}
	return nil
}
