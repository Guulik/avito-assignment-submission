package postgresql

import (
	"Avito_trainee_assignment/internal/domain/model"
	"Avito_trainee_assignment/internal/storage"
)

var _ storage.BannerStorage = (*Storage)(nil)

type Storage struct {
}

func New() *Storage {
	return &Storage{}
}

func (s Storage) UserBanner(featureId int, tagId int, lastRevision bool) (*model.BannerContent, error) {

	//TODO implement me
	return nil, storage.ErrNotFound

}

func (s Storage) Banners(featureId int, tagIg int, limit int, offset int) (*model.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) Save(featureId int, tagsId []int, content model.BannerContent, isActive bool) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) Patch(bannerId int, tagsId []int, featureId int, content model.BannerContent) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) Delete(bannerId int) error {
	//TODO implement me
	panic("implement me")
}
