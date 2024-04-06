package domain

import "Avito_trainee_assignment/internal/domain/model"

type GetUserRequest struct {
	FeatureId    int
	TagIg        int
	LastRevision bool
}

type GetRequest struct {
	FeatureId int
	TagIg     int
	Limit     int
	Offset    int
}

type CreateRequest struct {
	FeatureId int
	TagsId    []int
	Content   model.BannerContent
}

type UpdateRequest struct {
	BannerId  int
	TagsId    []int
	FeatureId int
	Content   model.BannerContent
}

type DeleteRequest struct {
	BannerId int
}
