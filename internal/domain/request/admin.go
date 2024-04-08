package request

type GetRequest struct {
	Token     string `header:"token"`
	FeatureId int    `query:"feature_id"`
	TagId     int    `query:"tag_id"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
}

type CreateRequest struct {
	Token     string                 `header:"token"`
	TagIds    []int                  `json:"tag_ids"`
	FeatureId int                    `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
}

type UpdateRequest struct {
	BannerId  int                    `param:"id"`
	Token     string                 `header:"token"`
	TagIds    []int                  `json:"tag_ids"`
	FeatureId int                    `json:"featureId"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
}

type DeleteRequest struct {
	BannerId int    `param:"id"`
	Token    string `header:"token"`
}
