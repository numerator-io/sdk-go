package response

type FeatureFlagResponseParams interface {
	FeatureFlag |
		FeatureFlagListResponse |
		FeatureFlagVariationValue
}

type UserOperationResponse[T FeatureFlagResponseParams] struct {
	Params T `json:"params"`
}
