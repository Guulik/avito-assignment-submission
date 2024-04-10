package request

type GetUserRequest struct {
	Token        string `header:"token"`
	FeatureId    int64  `query:"feature_id"`
	TagId        int64  `query:"tag_id"`
	LastRevision bool   `query:"use_last_revision"`
}
