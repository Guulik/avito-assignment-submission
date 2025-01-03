package api

import (
	"Banner_Infrastructure/internal/service"
	"log/slog"
)

type Api struct {
	log *slog.Logger
	svc service.BannerService
}

func New(
	log *slog.Logger,
	service service.BannerService,
) *Api {
	return &Api{
		log: log,
		svc: service,
	}
}
