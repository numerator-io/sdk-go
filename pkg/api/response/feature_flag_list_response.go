package response

type FeatureFlagListResponse struct {
	CountVal int64         `json:"count"`
	DataVal  []FeatureFlag `json:"data"`
}

func (r *FeatureFlagListResponse) Count() int64 {
	return r.CountVal
}

func (r *FeatureFlagListResponse) Data() []FeatureFlag {
	return r.DataVal
}
