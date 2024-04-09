package bannerSvc

import (
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"encoding/json"
	"log/slog"
)

func (s *Service) GetUserBanner(featureId int, tagId int, lastRevision bool) (map[string]interface{}, error) {
	const op = "Service.GetUserBanner"

	log := s.log.With(
		slog.String("op", op),
	)

	bannerJSON, err := s.storage.UserBannerDB(featureId, tagId)
	if err != nil {
		log.Error("failed to get bannerJSON for user", sl.Err(err))
		return nil, err
	}

	var content interface{}
	err = json.Unmarshal(bannerJSON, &content)
	if err != nil {
		log.Error("failed to unmarshal content from banner", sl.Err(err))
		return nil, err
	}

	return content.(map[string]interface{}), nil
}
