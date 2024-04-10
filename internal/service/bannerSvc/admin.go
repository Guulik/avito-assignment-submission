package bannerSvc

import (
	"Avito_trainee_assignment/internal/domain/model"
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
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

	SQLContent, err := json.Marshal(content)
	if err != nil {
		log.Error("failed to Marshal data", sl.Err(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	err = s.storage.Patch(bannerId, tagIds, featureId, SQLContent, isActive)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteBanner(bannerId int64) error {
	const op = "Service.UpdateBanner"

	_ = s.log.With(
		slog.String("op", op),
	)
	err := s.storage.Delete(bannerId)
	if err != nil {
		return err
	}

	return nil
}
