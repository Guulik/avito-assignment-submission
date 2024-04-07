package request

type GetUserRequest struct {
	Token        string `header:"token"`
	FeatureId    int    `query:"feature_id"`
	TagIg        int    `query:"tag_id"`
	LastRevision bool   `query:"use_last_revision"`
}
