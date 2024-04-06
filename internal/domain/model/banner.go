package model

import "time"

type Banner struct {
	featureId int32
	tagsId    []int32
	content   BannerContent
	isActive  bool
	createdAt time.Time
	updatedAt time.Time
}

type BannerContent struct {
}
