package banner

import (
	"Avito_trainee_assignment/internal/domain/model"
	"Avito_trainee_assignment/internal/repository"
)

var _ repository.BannerRepository = (*Repository)(nil)

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) DBGetUserBanner(featureId int, tagId int, lastRevision bool) *model.BannerContent {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) DBGetBanners(featureId int, tagIg int, limit int, offset int) *model.Banner {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) DBCreateBanner(featureId int, tagsId []int, content model.BannerContent) int {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) DBUpdateBanner(bannerId int, tagsId []int, featureId int, content model.BannerContent) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) DBDeleteBanner(bannerId int) error {
	//TODO implement me
	panic("implement me")
}
