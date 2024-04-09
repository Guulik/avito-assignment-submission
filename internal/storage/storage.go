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
		featureId int,
		tagId int,
	) ([]byte, error)
	UserBannerCached(
		featureId int,
		tagId int,
	) ([]byte, error)
	Banners(
		featureId int,
		tagIg int,
		limit int,
		offset int,
	) (*model.Banner, error)
	Save(
		featureId int,
		tagsId []int,
		content []byte,
		isActive bool,
	) (int, error)
	Patch(
		bannerId int,
		tagsId []int,
		featureId int,
		content []byte,
		isActive bool,
	) error
	Delete(bannerId int) error
}
