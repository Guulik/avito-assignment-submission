package api

import (
	"Avito_trainee_assignment/internal/domain/request"
	"Avito_trainee_assignment/internal/lib/binder"
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"Avito_trainee_assignment/internal/lib/validator"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) GetUserBanner(ctx echo.Context) error {
	const op = "Api.GetUserBanner"

	log := a.log.With(
		slog.String("op", op),
	)
	// default empty request values
	req := request.GetUserRequest{
		Token:        "",
		FeatureId:    -1,
		TagId:        -1,
		LastRevision: false,
	}

	// checks if request in correct form and bind it
	err := binder.BindReq(log, ctx, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	log.Info(sl.Req(req))

	if err = validator.Authorize(req.Token); err != nil {
		log.Error("incorrect token", sl.Err(err))
		return echo.NewHTTPError(http.StatusUnauthorized, "User unauthorized")
	}
	if err = validator.CheckGetRequest(&req); err != nil {
		log.Error("incorrect request", sl.Err(err))
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	banner, err := a.svc.GetUserBanner(
		req.FeatureId,
		req.TagId,
		req.LastRevision,
	)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, banner)
}
