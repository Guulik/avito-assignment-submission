package request

type GetRequest struct {
	Token     string `header:"token"`
	FeatureId int64  `query:"feature_id"`
	TagId     int64  `query:"tag_id"`
	Limit     int64  `query:"limit"`
	Offset    int64  `query:"offset"`
}

type CreateRequest struct {
	Token     string                 `header:"token"`
	TagIds    []int64                `json:"tag_ids"`
	FeatureId int64                  `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
}

type UpdateRequest struct {
	BannerId  int64                  `param:"id"`
	Token     string                 `header:"token"`
	TagIds    []int64                `json:"tag_ids"`
	FeatureId int64                  `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
}

type DeleteRequest struct {
	BannerId  int64  `param:"id"`
	FeatureId int64  `query:"feature_id"`
	TagId     int64  `query:"tag_id"`
	Token     string `header:"token"`
}
