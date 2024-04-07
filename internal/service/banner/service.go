package banner

import (
	"Avito_trainee_assignment/internal/service"
	"Avito_trainee_assignment/internal/storage"
	"log/slog"
)

var _ service.BannerService = (*Service)(nil)

type Service struct {
	log     *slog.Logger
	storage storage.BannerStorage
}

func New(
	log *slog.Logger,
	storage storage.BannerStorage,
) *Service {
	return &Service{
		log:     log,
		storage: storage,
	}
}
