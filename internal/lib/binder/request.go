package binder

import (
	sl "Banner_Infrastructure/internal/lib/logger/slog"
	"log/slog"

	"github.com/labstack/echo/v4"
)

func BindReq(log *slog.Logger, ctx echo.Context, req interface{}) error {
	err := ctx.Bind(req)
	if err != nil {
		log.Error("failed to parse query", sl.Err(err))
		return err
	}

	binder := &echo.DefaultBinder{}
	err = binder.BindHeaders(ctx, req)
	if err != nil {
		log.Error("failed to parse token", sl.Err(err))
		return err
	}

	return nil
}
