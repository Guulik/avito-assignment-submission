package banner

import (
	"Avito_trainee_assignment/internal/domain"
	"Avito_trainee_assignment/internal/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
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
	featureId = -1
	tagId = -1
	lastRevision = false
	token = ""

	return &Api{
		svc: service,
	}
}

func (a *Api) GetUserBanner(ctx echo.Context) error {
	var request domain.GetUserRequest
	err := ctx.Bind(&request)
	if err != nil {
		fmt.Println(err)
	}

	binder := &echo.DefaultBinder{}
	err = binder.BindHeaders(ctx, request)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", request)

	/*if !(validator.Validate(request.Token) == validator.User) {
		return echo.NewHTTPError(http.StatusForbidden)
	}*/

	_, err = a.svc.GetUserBanner(request)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	//TODO implement me

	return nil
}
