package model

import "time"

type Banner struct {
	ID        int32
	FeatureId int32
	TagIds    []int32
	Content   map[string]interface{}
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
