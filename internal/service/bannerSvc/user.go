package bannerSvc

import (
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"fmt"
	"log/slog"
)

func (s *Service) GetUserBanner(featureId int, tagId int, lastRevision bool) (map[string]interface{}, error) {
	const op = "Service.GetUserBanner"

	log := s.log.With(
		slog.String("op", op),
	)

	banner, err := s.storage.UserBanner(featureId, tagId, lastRevision)
	if err != nil {
		log.Error("failed to get banner for user", sl.Err(err))
		return nil, err
	}

	//test_banner := model.Banner{ID: 1, FeatureId: 45, TagIds: []int32{5, 12, 244}}
	if !banner.IsActive {
		log.Warn("attempt to get inactive banner")
		return nil, fmt.Errorf("banner {id:%d, featureId:%d, tagIds:%v} is inactive",
			banner.ID, banner.FeatureId, banner.TagIds)
	}

	return banner.Content, nil
}
