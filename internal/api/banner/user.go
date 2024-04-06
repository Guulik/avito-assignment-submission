package banner

import (
	"Avito_trainee_assignment/internal/domain"
	"Avito_trainee_assignment/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var (
	featureId    int
	tagId        int
	lastRevision bool
	token        string
)

type Api struct {
	svc service.BannerService
}

func New(service service.BannerService) *Api {
	return &Api{
		svc: service,
	}
}

func (a *Api) GetUserBanner(ctx echo.Context) error {
	featureId, _ = strconv.Atoi(ctx.QueryParam("feature_id"))
	if !validToken(token) {
		//TODO return forbidden 403
	}

	request := domain.GetUserRequest{}
	_, err := a.svc.GetUserBanner(request)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	//TODO implement me

	return nil
}

func validToken(token string) bool {

	panic("implement me")
}
