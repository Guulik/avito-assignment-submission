package bannerSvc

import (
	"Avito_trainee_assignment/internal/domain/model"
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"encoding/json"
	"log/slog"
)

func (s *Service) GetBanners(featureId int, tagId int, limit int, offset int) (*model.Banner, error) {
	const op = "Service.GetBanners"

	log := s.log.With(
		slog.String("op", op),
	)
	_, err := s.storage.Banners(featureId, tagId, limit, offset)
	if err != nil {
		log.Error("failed to get banners list", sl.Err(err))
		return nil, err
	}
	//TODO: return banners instead of nil
	return nil, nil
}

func (s *Service) CreateBanner(featureId int, tagIds []int, content map[string]interface{}, isActive bool) (int, error) {
	const op = "Service.CreateBanner"

	log := s.log.With(
		slog.String("op", op),
	)
	SQLContent, err := json.Marshal(content)
	if err != nil {
		log.Error("failed to Marshal data", sl.Err(err))
		return -1, err
	}

	bannerId, err := s.storage.Save(featureId, tagIds, SQLContent, isActive)
	if err != nil {
		log.Error("failed to create banner", sl.Err(err))
		return -1, err
	}
	return bannerId, nil
}

func (s *Service) UpdateBanner(bannerId int, tagIds []int, featureId int, content map[string]interface{}, isActive bool) error {
	const op = "Service.UpdateBanner"

	log := s.log.With(
		slog.String("op", op),
	)

	SQLContent, err := json.Marshal(content)
	if err != nil {
		log.Error("failed to Marshal data", sl.Err(err))
		return err
	}

	err = s.storage.Patch(bannerId, tagIds, featureId, SQLContent, isActive)
	if err != nil {
		log.Error("failed to update banner", sl.Err(err))
		return err
	}

	return nil
}

func (s *Service) DeleteBanner(bannerId int) error {
	const op = "Service.UpdateBanner"

	log := s.log.With(
		slog.String("op", op),
	)
	err := s.storage.Delete(bannerId)
	if err != nil {
		log.Error("failed to delete banner", sl.Err(err))
		return err
	}

	return nil
}
