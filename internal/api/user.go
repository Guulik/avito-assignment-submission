package api

import (
	"Avito_trainee_assignment/internal/domain/request"
	"Avito_trainee_assignment/internal/lib/binder"
	"Avito_trainee_assignment/internal/lib/jwt/validator"
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

func (a *Api) GetUserBanner(ctx echo.Context) error {
	const op = "Api.GetUserBanner"

	log := a.log.With(
		slog.String("op", op),
	)

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

	if err = validator.CheckUser(req.Token); err != nil {
		a.log.Error("incorrect token", sl.Err(err))
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	_, err = a.svc.GetUserBanner(
		req.TagId,
		req.FeatureId,
		req.LastRevision,
	)

	if err != nil {
		return ctx.String(http.StatusOK, err.Error())
	}

	return ctx.String(http.StatusOK, "its not a banner it is dummy response")
}
