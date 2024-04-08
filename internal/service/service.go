package service

import (
	"Avito_trainee_assignment/internal/domain/model"
)

type BannerService interface {
	GetUserBanner(
		featureId int,
		tagId int,
		lastRevision bool,
	) (map[string]interface{}, error)
	GetBanners(
		featureId int,
		tagId int,
		limit int,
		offset int,
	) (*model.Banner, error)
	CreateBanner(
		featureId int,
		tagIds []int,
		content map[string]interface{},
		isActive bool,
	) (int, error)
	UpdateBanner(
		bannerId int,
		tagIds []int,
		featureId int,
		content map[string]interface{},
		isActive bool,
	) error
	DeleteBanner(
		bannerId int,
	) error
}
