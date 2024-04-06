package banner

import (
	"Avito_trainee_assignment/internal/domain"
	"Avito_trainee_assignment/internal/domain/model"
	"Avito_trainee_assignment/internal/repository"
	"Avito_trainee_assignment/internal/service"
	"errors"
)

var _ service.BannerService = (*Service)(nil)

type Service struct {
	repo repository.BannerRepository
}

func New(repository repository.BannerRepository) *Service {
	return &Service{
		repo: repository,
	}
}

func (s *Service) GetUserBanner(req domain.GetUserRequest) (*model.BannerContent, error) {
	//TODO удалить этот ретурн
	return nil, errors.New("НЕТ ТАКОГО БАННЕРА И ВООБЩЕ НИЧЕГО ПОКА НЕ РЕАЛИЗОВАНО")

	_ = s.repo.DBGetUserBanner(req.FeatureId, req.TagIg, req.LastRevision)
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetBanners(req domain.GetRequest) (*model.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CreateBanner(req domain.CreateRequest) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateBanner(req domain.UpdateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) DeleteBanner(req domain.DeleteRequest) error {
	//TODO implement me
	panic("implement me")
}
