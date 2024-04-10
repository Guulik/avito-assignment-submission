package api

import (
	"Avito_trainee_assignment/internal/domain/request"
	"Avito_trainee_assignment/internal/lib/binder"
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"Avito_trainee_assignment/internal/lib/validator"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

func (a *Api) GetUserBanner(ctx echo.Context) error {
	const op = "Api.GetUserBanner"

	log := a.log.With(
		slog.String("op", op),
	)
	//default empty request values
	req := request.GetUserRequest{
		Token:        "",
		FeatureId:    -1,
		TagId:        -1,
		LastRevision: false,
	}

	err := binder.BindReq(log, ctx, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	log.Info(sl.Req(req))

	if err = validator.Authorize(req.Token); err != nil {
		log.Error("incorrect token", sl.Err(err))
		return echo.NewHTTPError(http.StatusUnauthorized, "Пользователь не авторизован")
	}
	if err = validator.CheckGetUserRequest(req); err != nil {
		log.Error("incorrect request", sl.Err(err))
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	banner, err := a.svc.GetUserBanner(
		req.FeatureId,
		req.TagId,
		req.LastRevision,
	)
	if err != nil {
		log.Error("failed to get banner for user", sl.Err(err))
		return err
	}

	return ctx.JSON(http.StatusOK, banner)
}
