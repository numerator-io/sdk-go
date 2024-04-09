package response

type FeatureFlagList struct {
	CountVal int64         `json:"count"`
	DataVal  []FeatureFlag `json:"data"`
}

func (r *FeatureFlagList) Count() int64 {
	return r.CountVal
}

func (r *FeatureFlagList) Data() []FeatureFlag {
	return r.DataVal
}
