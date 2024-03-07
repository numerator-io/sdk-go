package clients

import (
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/pkg/config"
)

type NumeratorClient interface {
	FeatureFlags(page, size int) ([]response.FeatureFlag, error)
	FeatureFlagDetails(flagKey string) (*response.FeatureFlag, error)
	GetValueByKeyWithDefault(flagKey string, context map[string]interface{}, defaultValue interface{}) (interface{}, error)
}

func NewNumeratorClient(apiKey string, numeratorConfig *config.NumeratorConfig) NumeratorClient {
	if numeratorConfig == nil {
		numeratorConfig = &config.NumeratorConfig{APIKey: apiKey}
	} else {
		numeratorConfig.APIKey = apiKey
	}
	return NewDefaultNumeratorClient(numeratorConfig)
}
