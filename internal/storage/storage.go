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
	UserBanner(
		featureId int,
		tagId int,
		lastRevision bool,
	) (*model.Banner, error)
	Banners(
		featureId int,
		tagIg int,
		limit int,
		offset int,
	) (*model.Banner, error)
	Save(
		featureId int,
		tagsId []int,
		content map[string]interface{},
		isActive bool,
	) (int, error)
	Patch(
		bannerId int,
		tagsId []int,
		featureId int,
		content map[string]interface{},
		isActive bool,
	) error
	Delete(bannerId int) error
}
