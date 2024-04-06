package service

import (
	"Avito_trainee_assignment/internal/domain"
	"Avito_trainee_assignment/internal/domain/model"
)

type BannerService interface {
	GetUserBanner(request domain.GetUserRequest) (*model.BannerContent, error)
	GetBanners(request domain.GetRequest) (*model.Banner, error)
	CreateBanner(request domain.CreateRequest) (int, error)
	UpdateBanner(request domain.UpdateRequest) error
	DeleteBanner(request domain.DeleteRequest) error
}
