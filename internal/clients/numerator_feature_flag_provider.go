package clients

import (
	"fmt"
	"net/http"

	"github.com/numerator-io/sdk-go/pkg/api/response"
	"github.com/numerator-io/sdk-go/pkg/config"
	"github.com/numerator-io/sdk-go/pkg/context"
	"github.com/numerator-io/sdk-go/pkg/exception"
)

type NumeratorFeatureFlagProvider struct {
	numeratorConfig *config.NumeratorConfig
	contextProvider context.ContextProvider
	client          NumeratorClient
}

func NewNumeratorFeatureFlagProvider(nConfig *config.NumeratorConfig, contextProvider context.ContextProvider) *NumeratorFeatureFlagProvider {
	if contextProvider == nil {
		contextProvider = context.NewContextProvider()
	}

	provider := &NumeratorFeatureFlagProvider{
		numeratorConfig: nConfig,
		contextProvider: contextProvider,
		client:          NewNumeratorClient(nConfig, contextProvider),
	}
	return provider
}

func (p *NumeratorFeatureFlagProvider) GetBooleanFeatureFlag(key string, defaultValue bool, context map[string]interface{}, useDefaultContext bool) bool {
	val, err := p.getFlagValue(key, defaultValue, context, useDefaultContext)
	if err != nil {
		return defaultValue
	}
	result, ok := val.(*response.FlagEvaluationDetail[bool])
	if !ok {
		return defaultValue
	}
	return result.Value
}

func (p *NumeratorFeatureFlagProvider) GetStringFeatureFlag(key string, defaultValue string, context map[string]interface{}, useDefaultContext bool) string {
	val, err := p.getFlagValue(key, defaultValue, context, useDefaultContext)
	if err != nil {
		return defaultValue
	}
	result, ok := val.(*response.FlagEvaluationDetail[string])
	if !ok {
		return defaultValue
	}
	return result.Value
}

func (p *NumeratorFeatureFlagProvider) GetLongFeatureFlag(key string, defaultValue int64, context map[string]interface{}, useDefaultContext bool) int64 {
	val, err := p.getFlagValue(key, defaultValue, context, useDefaultContext)
	if err != nil {
		return defaultValue
	}
	result, ok := val.(*response.FlagEvaluationDetail[int64])
	if !ok {
		return defaultValue
	}
	return result.Value
}

func (p *NumeratorFeatureFlagProvider) GetDoubleFeatureFlag(key string, defaultValue float64, context map[string]interface{}, useDefaultContext bool) float64 {
	val, err := p.getFlagValue(key, defaultValue, context, useDefaultContext)
	if err != nil {
		return defaultValue
	}
	result, ok := val.(*response.FlagEvaluationDetail[float64])
	if !ok {
		return defaultValue
	}
	return result.Value
}

func (p *NumeratorFeatureFlagProvider) getFlagValue(key string, defaultValue interface{}, context map[string]interface{}, useDefaultContext bool) (interface{}, error) {
	switch defaultValue := defaultValue.(type) {
	case bool:
		return p.client.BooleanFlagVariationDetail(key, context, defaultValue, useDefaultContext)
	case string:
		return p.client.StringFlagVariationDetail(key, context, defaultValue, useDefaultContext)
	case int64:
		return p.client.LongFlagVariationDetail(key, context, defaultValue, useDefaultContext)
	case float64:
		return p.client.DoubleFlagVariationDetail(key, context, defaultValue, useDefaultContext)
	default:
		return nil, exception.NewNumeratorException(fmt.Sprintf("Unsupported flag type %T", defaultValue), http.StatusInternalServerError)
	}
}
