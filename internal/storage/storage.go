package storage

import (
	"Avito_trainee_assignment/internal/domain/model"
	"errors"
)

var (
	ErrNotFound = errors.New("banner not found in DB")
)

type BannerStorage interface {
	UserBanner(
		featureId int,
		tagId int,
		lastRevision bool,
	) (*model.BannerContent, error)
	Banners(
		featureId int,
		tagIg int,
		limit int,
		offset int,
	) (*model.Banner, error)
	Save(
		featureId int,
		tagsId []int,
		content model.BannerContent,
		isActive bool,
	) (int, error)
	Patch(
		bannerId int,
		tagsId []int,
		featureId int,
		content model.BannerContent,
	) error
	Delete(bannerId int) error
}
