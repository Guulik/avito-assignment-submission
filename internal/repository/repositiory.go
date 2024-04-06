package repository

import (
	"Avito_trainee_assignment/internal/domain/model"
)

type BannerRepository interface {
	DBGetUserBanner(featureId int, tagId int, lastRevision bool) *model.BannerContent
	DBGetBanners(featureId int, tagIg int, limit int, offset int) *model.Banner
	DBCreateBanner(featureId int, tagsId []int, content model.BannerContent) int
	DBUpdateBanner(bannerId int, tagsId []int, featureId int, content model.BannerContent) error
	DBDeleteBanner(bannerId int) error
}
