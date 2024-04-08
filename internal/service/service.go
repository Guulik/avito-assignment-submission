package service

import (
	"Avito_trainee_assignment/internal/domain/model"
)

type BannerService interface {
	GetUserBanner(
		featureId int,
		tagId int,
		lastRevision bool,
	) (*model.BannerContent, error)
	GetBanners(
		featureId int,
		tagIg int,
		limit int,
		offset int,
	) (*model.Banner, error)
	CreateBanner(
		featureId int,
		tagsId []int,
		content model.BannerContent,
		isActive bool,
	) (int, error)
	UpdateBanner(
		bannerId int,
		tagsId []int,
		featureId int,
		content model.BannerContent,
		isActive bool,
	) error
	DeleteBanner(
		bannerId int,
	) error
}
