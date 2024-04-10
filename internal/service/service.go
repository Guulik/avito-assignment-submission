package service

import (
	"Avito_trainee_assignment/internal/domain/model"
)

type BannerService interface {
	GetUserBanner(
		featureId int64,
		tagId int64,
		lastRevision bool,
	) (map[string]interface{}, error)
	GetBanners(
		featureId int64,
		tagId int64,
		limit int64,
		offset int64,
	) ([]model.Banner, error)
	CreateBanner(
		featureId int64,
		tagIds []int64,
		content map[string]interface{},
		isActive bool,
	) (int64, error)
	UpdateBanner(
		bannerId int64,
		tagIds []int64,
		featureId int64,
		content map[string]interface{},
		isActive bool,
	) error
	DeleteBanner(
		bannerId int64,
	) error
}
