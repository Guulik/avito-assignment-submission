package bannerSvc

import (
	sl "Banner_Infrastructure/internal/lib/logger/slog"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetUserBanner(featureId int64, tagId int64, lastRevision bool) (map[string]interface{}, error) {
	const op = "Service.GetUserBanner"

	log := s.log.With(
		slog.String("op", op),
	)
	var (
		bannerJSON []byte
		err        error
	)
	if lastRevision {
		log.Info("getting from DB")
		bannerJSON, err = s.storage.UserBannerDB(featureId, tagId)
		if err != nil {
			return nil, err
		}
	} else {
		log.Info("getting from cache")
		bannerJSON, err = s.cache.GetBannerCached(featureId, tagId)
		if err != nil {
			bannerJSON, err = s.storage.UserBannerDB(featureId, tagId)
			if err != nil {
				return nil, err
			}
			err = s.cache.SetBannerCache(featureId, tagId, bannerJSON)
			if err != nil {
				log.Error(sl.Err(err).String())
			}
		}
	}

	var content map[string]interface{}
	err = json.Unmarshal(bannerJSON, &content)
	if err != nil {
		log.Error("failed to unmarshal content from banner", sl.Err(err))
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return content, nil
}
