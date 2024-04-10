package bannerSvc

import (
	"Avito_trainee_assignment/internal/service"
	"Avito_trainee_assignment/internal/storage"
	"log/slog"
)

var _ service.BannerService = (*Service)(nil)

type Service struct {
	log     *slog.Logger
	storage storage.BannerStorage
	cache   storage.BannerCache
}

func New(
	log *slog.Logger,
	storage storage.BannerStorage,
	cache storage.BannerCache,
) *Service {
	return &Service{
		log:     log,
		storage: storage,
		cache:   cache,
	}
}
