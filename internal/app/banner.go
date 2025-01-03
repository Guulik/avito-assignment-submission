package app

import (
	"Banner_Infrastructure/internal/api"
	"Banner_Infrastructure/internal/configure"
	sl "Banner_Infrastructure/internal/lib/logger/slog"
	service "Banner_Infrastructure/internal/service/bannerSvc"
	"Banner_Infrastructure/internal/storage/postgresql"
	"Banner_Infrastructure/internal/storage/redis"
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type App struct {
	api     *api.Api
	svc     *service.Service
	storage *postgresql.Storage
	echo    *echo.Echo
	pool    *sqlx.DB
	cache   *redis.Cache
}

func New(log *slog.Logger, cfg *configure.Config) *App {
	app := &App{}

	app.echo = echo.New()

	app.pool = configure.NewPostgres(cfg)

	rds := redis.InitRedis(cfg)

	app.storage = postgresql.New(log, app.pool)

	app.cache = redis.New(log, rds, cfg)

	app.svc = service.New(log, app.storage, app.cache)

	app.api = api.New(log, app.svc)

	err := cfg.MigrateUp()
	if err != nil {
		log.Error("failed to connect to create table in DB", sl.Err(err))
	}

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
	if err := a.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

func (a *App) Stop(ctx context.Context) error {
	fmt.Println("stopping server..." + " op = app.Stop")

	if err := a.echo.Shutdown(ctx); err != nil {
		fmt.Println("failed to shutdown server")
		return err
	}

	if err := a.pool.Close(); err != nil {
		fmt.Println("failed to close connection")
	}
	return nil
}
