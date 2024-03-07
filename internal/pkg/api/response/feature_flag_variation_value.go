package response

import (
	"github.com/c0x12c/numerator-go-sdk/internal/pkg/enum"
)

type FeatureFlagVariationValue struct {
	Key       string             `json:"key"`
	Status    string             `json:"status"`
	Value     VariationValue     `json:"value"`
	ValueType enum.FlagValueType `json:"value_type"`
}
