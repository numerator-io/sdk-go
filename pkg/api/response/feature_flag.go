package response

import (
	"time"

	"github.com/c0x12c/numerator-go-sdk/pkg/enum"
)

type FeatureFlag struct {
	Name                string             `json:"name"`
	Key                 string             `json:"key"`
	Status              string             `json:"status"`
	Description         *string            `json:"description,omitempty"`
	DefaultOnVariation  FlagVariation      `json:"default_on_variation"`
	DefaultOffVariation FlagVariation      `json:"default_off_variation"`
	ValueType           enum.FlagValueType `json:"value_type"` // Use the enum.FlagValueType type
	CreatedAt           time.Time          `json:"created_at"`
}

type FlagVariation struct {
	Name  string         `json:"name"`
	Value VariationValue `json:"value"`
}
