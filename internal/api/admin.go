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

func (a *Api) GetBanner(ctx echo.Context) error {
	const op = "Api.GetBanner"

	log := a.log.With(
		slog.String("op", op),
	)

	req := request.GetRequest{
		Token:     "",
		FeatureId: -1,
		TagId:     -1,
		Limit:     -1,
		Offset:    -1,
	}

	err := binder.BindReq(log, ctx, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	log.Info(sl.Req(req))

	if err = validator.CheckAdmin(req.Token); err != nil {
		a.log.Error("incorrect token", sl.Err(err))
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	_, err = a.svc.GetBanners(
		req.FeatureId,
		req.TagId,
		req.Limit,
		req.Offset,
	)

	if err != nil {
		return ctx.String(http.StatusOK, err.Error())
	}
	return ctx.String(http.StatusOK, "its not a banner it is dummy response")
}

func (a *Api) CreateBanner(ctx echo.Context) error {
	const op = "Api.CreateBanner"

	log := a.log.With(
		slog.String("op", op),
	)

	req := request.CreateRequest{
		Token:     "",
		TagIds:    nil,
		FeatureId: -1,
		Content:   nil,
		IsActive:  false,
	}

	err := binder.BindReq(log, ctx, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	log.Info(sl.Req(req))

	if err = validator.CheckAdmin(req.Token); err != nil {
		a.log.Error("incorrect token", sl.Err(err))
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	_, err = a.svc.CreateBanner(
		req.FeatureId,
		req.TagIds,
		req.Content,
		req.IsActive,
	)

	if err != nil {
		return ctx.String(http.StatusOK, err.Error())
	}
	return ctx.String(http.StatusOK, "its not a banner it is dummy response")
}

func (a *Api) PatchBanner(ctx echo.Context) error {
	const op = "Api.PatchBanner"

	log := a.log.With(
		slog.String("op", op),
	)

	req := request.UpdateRequest{
		BannerId:  -1,
		Token:     "",
		TagIds:    nil,
		FeatureId: -1,
		Content:   nil,
		IsActive:  false,
	}

	err := binder.BindReq(log, ctx, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	log.Info(sl.Req(req))

	if err = validator.CheckAdmin(req.Token); err != nil {
		a.log.Error("incorrect token", sl.Err(err))
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	err = a.svc.UpdateBanner(
		req.BannerId,
		req.TagIds,
		req.FeatureId,
		req.Content,
		req.IsActive,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.String(http.StatusOK, "its not a banner it is dummy response")
}

func (a *Api) DeleteBanner(ctx echo.Context) error {
	const op = "Api.DeleteBanner"

	log := a.log.With(
		slog.String("op", op),
	)

	req := request.DeleteRequest{
		BannerId: -1,
		Token:    "",
	}

	err := binder.BindReq(log, ctx, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	log.Info(sl.Req(req))

	if err = validator.CheckAdmin(req.Token); err != nil {
		a.log.Error("incorrect token", sl.Err(err))
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	err = a.svc.DeleteBanner(
		req.BannerId,
	)

	return ctx.String(http.StatusOK, "its not a banner it is dummy response")
}
