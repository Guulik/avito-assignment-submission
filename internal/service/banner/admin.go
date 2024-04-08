package banner

import "Avito_trainee_assignment/internal/domain/model"

func (s *Service) GetBanners(featureId int, tagIg int, limit int, offset int) (*model.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CreateBanner(featureId int, tagsId []int, content model.BannerContent, isActive bool) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateBanner(bannerId int, tagsId []int, featureId int, content model.BannerContent, isActive bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) DeleteBanner(bannerId int) error {
	//TODO implement me
	panic("implement me")
}
