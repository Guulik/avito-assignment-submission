package bannerSvc

import (
	"Banner_Infrastructure/internal/domain/model"
	sl "Banner_Infrastructure/internal/lib/logger/slog"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetBanners(featureId int64, tagId int64, limit int64, offset int64) ([]model.Banner, error) {
	const op = "Service.GetBanners"

	log := s.log.With(
		slog.String("op", op),
	)
	var (
		bannersDb []model.BannerDB
		err       error
	)

	if featureId != -1 || tagId != -1 {
		bannersDb, err = s.storage.FilteredBanners(featureId, tagId, limit, offset)
		if err != nil {
			return nil, err
		}
	} else {
		bannersDb, err = s.storage.Banners(limit, offset)
		if err != nil {
			return nil, err
		}
	}

	banners, err := model.BannersDbToService(bannersDb)
	if err != nil {
		log.Error("failed to convert db banner to service model", sl.Err(err))
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return banners, nil
}

func (s *Service) CreateBanner(featureId int64, tagIds []int64, content map[string]interface{}, isActive bool) (int64, error) {
	const op = "Service.CreateBanner"

	log := s.log.With(
		slog.String("op", op),
	)
	SQLContent, err := json.Marshal(content)
	if err != nil {
		log.Error("failed to Marshal data", sl.Err(err))
		return -1, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	bannerId, err := s.storage.Save(featureId, tagIds, SQLContent, isActive)
	if err != nil {
		return -1, err
	}
	return bannerId, nil
}

func (s *Service) UpdateBanner(bannerId int64, tagIds []int64, featureId int64, content map[string]interface{}, isActive bool) error {
	const op = "Service.UpdateBanner"

	log := s.log.With(
		slog.String("op", op),
	)

	cleanCahce := false
	if len(tagIds) > 0 || featureId > 0 {
		cleanCahce = true
	}

	currentBannerDb, err := s.storage.GetBannerById(bannerId)
	if err != nil {
		log.Error("failed to get banner by Id", sl.Err(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	currentBanner, err := model.ToBanner(*currentBannerDb)
	if err != nil {
		log.Error("failed to convert banner to service", sl.Err(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if len(tagIds) == 0 {
		tagIds = currentBanner.TagIds
	}
	if featureId < 0 {
		featureId = currentBanner.FeatureId
	}
	if content == nil {
		content = currentBanner.Content
	}
	if isActive != currentBanner.IsActive {
		isActive = currentBanner.IsActive
	}
	SQLContent, err := json.Marshal(content)
	if err != nil {
		log.Error("failed to Marshal data", sl.Err(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cleanCahce {
		bannerDb, _ := s.storage.GetBannerById(bannerId)
		banner, _ := model.ToBanner(*bannerDb)
		for _, tag := range banner.TagIds {
			err = s.cache.DeleteBannerCache(banner.FeatureId, tag)
		}
		if err != nil {
			log.Warn("failed to clean cache", sl.Err(err))
		}
	}

	err = s.storage.Patch(bannerId, tagIds, featureId, SQLContent, isActive)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteBanner(bannerId int64, featureId int64, tagId int64) error {
	const op = "Service.UpdateBanner"

	log := s.log.With(
		slog.String("op", op),
	)
	err := s.storage.Delete(bannerId, featureId, tagId)
	if err != nil {
		return err
	}

	log.Debug("deleting from cache")
	err = s.cache.DeleteBannerCache(featureId, tagId)
	if err != nil {
		log.Warn("failed to delete from cache")
	}

	return nil
}
