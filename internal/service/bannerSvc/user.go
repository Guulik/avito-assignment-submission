package bannerSvc

import (
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

func (s *Service) GetUserBanner(featureId int64, tagId int64, lastRevision bool) (map[string]interface{}, error) {
	const op = "Service.GetUserBanner"

	log := s.log.With(
		slog.String("op", op),
	)

	bannerJSON, err := s.storage.UserBannerDB(featureId, tagId)
	if err != nil {
		return nil, err
	}

	var content interface{}
	err = json.Unmarshal(bannerJSON, &content)
	if err != nil {
		log.Error("failed to unmarshal content from banner", sl.Err(err))
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return content.(map[string]interface{}), nil
}
