package model

import "time"

type Banner struct {
	ID        int64                  `json:"banner_id"`
	FeatureId int64                  `json:"feature_id"`
	TagIds    []int64                `json:"tag_ids"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type BannerDB struct {
	ID        int64     `db:"banner_id"`
	FeatureId int64     `db:"feature_id"`
	TagIds    string    `db:"tag_ids"`
	Content   []byte    `db:"content"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
