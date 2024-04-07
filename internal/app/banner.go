package app

import (
	api "Avito_trainee_assignment/internal/api/banner"
	repo "Avito_trainee_assignment/internal/repository/banner"
	service "Avito_trainee_assignment/internal/service/banner"
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
)

type App struct {
	api  *api.Api
	svc  *service.Service
	repo *repo.Repository
	echo *echo.Echo
}

func New(log *slog.Logger) *App {
	app := &App{}

	app.echo = echo.New()

	app.repo = repo.New()

	app.svc = service.New(app.repo)

	app.api = api.New(app.svc)

	//app.echo.POST("/Create", app.api.Create)
	app.echo.GET("/user_banner/get", app.api.GetUserBanner)
	//app.echo.GET("/Get", app.api.Get)
	//app.echo.PATCH("/Get", app.api.Get)
	//app.echo.DELETE("/Get", app.api.Get)

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
