package banner

import (
	"Avito_trainee_assignment/internal/domain/model"
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"log/slog"
)

func (s *Service) GetUserBanner(featureId int, tagId int, lastRevision bool) (*model.BannerContent, error) {
	const op = "Service.GetUserBanner"

	log := s.log.With(
		slog.String("op", op),
	)
	_, err := s.storage.UserBanner(featureId, tagId, lastRevision)
	if err != nil {
		log.Error("failed to get banner for user", sl.Err(err))
		return nil, err
	}

	//TODO implement me
	return nil, nil
}
