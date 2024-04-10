package storage

import (
	"Avito_trainee_assignment/internal/domain/model"
	"errors"
)

var (
	ErrNotFound   = errors.New("banner not found in DB")
	ErrSaveFail   = errors.New("failed to write banner to DB")
	ErrDeleteFail = errors.New("failed to remove banner from DB")
)

type BannerStorage interface {
	UserBannerDB(
		featureId int64,
		tagId int64,
	) ([]byte, error)
	Banners(
		limit int64,
		offset int64,
	) ([]model.BannerDB, error)
	FilteredBanners(
		featureId int64,
		tagIg int64,
		limit int64,
		offset int64,
	) ([]model.BannerDB, error)
	Save(
		featureId int64,
		tagsId []int64,
		content []byte,
		isActive bool,
	) (int64, error)
	Patch(
		bannerId int64,
		tagsId []int64,
		featureId int64,
		content []byte,
		isActive bool,
	) error
	Delete(bannerId int64) error
}

type BannerCache interface {
	GetBannerCached(
		featureId int64,
		tagId int64,
	) ([]byte, error)
	SetBannerCache(
		featureId int64,
		tagId int64,
		content []byte,
	) error
}
