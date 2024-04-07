package domain

import "Avito_trainee_assignment/internal/domain/model"

type GetUserRequest struct {
	FeatureId    int  `query:"feature_id"`
	TagIg        int  `query:"tag_id"`
	LastRevision bool `query:"use_last_revision"`
	//Token        string `header:"token"`
}

type GetRequest struct {
	FeatureId int    `query:"feature_id"`
	TagIg     int    `query:"tag_id"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
	Token     string `header:"token"`
}

type CreateRequest struct {
	Token     string              `header:"token"`
	FeatureId int                 `json:"feature_id"`
	TagIds    []int               `json:"tag_ids"`
	Content   model.BannerContent `json:"content"`
	IsActive  bool                `json:"is_active"`
}

type UpdateRequest struct {
	BannerId  int
	TagsId    []int
	FeatureId int
	Content   model.BannerContent
	Token     string `header:"token"`
}

type DeleteRequest struct {
	BannerId int
	Token    string `header:"token"`
}
