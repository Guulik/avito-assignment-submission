package api

import (
	"Avito_trainee_assignment/internal/domain/request"
	"Avito_trainee_assignment/internal/lib/binder"
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"Avito_trainee_assignment/internal/lib/validator"
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

func (a *Api) GetBanner(ctx echo.Context) error {
	const op = "Api.GetBanner"

	log := a.log.With(
		slog.String("op", op),
	)
	//default empty request values
	req := request.GetRequest{
		Token:     "",
		FeatureId: -1,
		TagId:     -1,
		Limit:     -1,
		Offset:    -1,
	}

	err := binder.BindReq(log, ctx, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	log.Info(sl.Req(req))

	if err = validator.CheckAdmin(req.Token); err != nil {
		a.log.Error("incorrect token", sl.Err(err))
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	banners, err := a.svc.GetBanners(
		req.FeatureId,
		req.TagId,
		req.Limit,
		req.Offset,
	)

	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, banners)
}

func (a *Api) CreateBanner(ctx echo.Context) error {
	const op = "Api.CreateBanner"

	log := a.log.With(
		slog.String("op", op),
	)
	//default empty request values
	req := request.CreateRequest{
		Token:     "",
		TagIds:    nil,
		FeatureId: -1,
		Content:   nil,
		IsActive:  true,
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
	if err = validator.CheckPostRequest(&req, true); err != nil {
		a.log.Error("incorrect request", sl.Err(err))
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	bannerId, err := a.svc.CreateBanner(
		req.FeatureId,
		req.TagIds,
		req.Content,
		req.IsActive,
	)

	if err != nil {
		return err
	}

	return ctx.String(http.StatusCreated, fmt.Sprintf("created banner with ID = %v", bannerId))
}

func (a *Api) PatchBanner(ctx echo.Context) error {
	const op = "Api.PatchBanner"

	log := a.log.With(
		slog.String("op", op),
	)
	//default empty request values
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
	if err = validator.CheckUpdateRequest(&req, true); err != nil {
		a.log.Error("incorrect request", sl.Err(err))
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = a.svc.UpdateBanner(
		req.BannerId,
		req.TagIds,
		req.FeatureId,
		req.Content,
		req.IsActive,
	)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("banner with Id %v was successfully updated", req.BannerId))
}

func (a *Api) DeleteBanner(ctx echo.Context) error {
	const op = "Api.DeleteBanner"

	log := a.log.With(
		slog.String("op", op),
	)
	//default empty request values
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
	if err = validator.CheckDeleteRequest(&req); err != nil {
		a.log.Error("incorrect request", sl.Err(err))
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = a.svc.DeleteBanner(
		req.BannerId,
	)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusNoContent, fmt.Sprintf("banner with Id %v was successfully deleted", req.BannerId))
}
