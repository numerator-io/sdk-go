package response

import (
	"github.com/numerator-io/sdk-go/pkg/enum"
)

type FeatureFlagVariationValue struct {
	Key       string             `json:"key"`
	Status    string             `json:"status"`
	Value     VariationValue     `json:"value"`
	ValueType enum.FlagValueType `json:"value_type"`
}
