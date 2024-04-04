//go:generate mockgen -source numerator_client.go -destination ./mock_client/mock_numerator_client/numerator_client_gen.go
package clients

import (
	"github.com/c0x12c/numerator-go-sdk/internal/service"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/pkg/config"
	"github.com/c0x12c/numerator-go-sdk/pkg/context"
	"github.com/c0x12c/numerator-go-sdk/pkg/network"
)

/**
 * Numerator client interface
 */
type NumeratorClient interface {
	/**
	 * Get the context provider
	 * Context provider is used to provide context conditions for feature flag evaluation
	 */
	GetContextProvider() context.ContextProvider

	/**
	 * Set the context provider
	 * Context provider is used to provide context conditions for feature flag evaluation
	 */
	SetContextProvider(contextProvider context.ContextProvider)

	/**
	 * Get the version of the SDK client
	 */
	Version() string

	/**
	 * Get a list of feature flags in the current project.
	 *
	 * @param page page number
	 * @param size number of results per page
	 * @return a list of feature flags
	 */
	FeatureFlags(page, size int) ([]response.FeatureFlag, error)

	/**
	 * Get details of a single feature flag.
	 *
	 * @param flagKey feature flag identifying key
	 * @return details of feature flag if found
	 */
	FeatureFlagDetails(flagKey string) (*response.FeatureFlag, error)

	/**
	 * Get the Variation of the feature flag.
	 *
	 * @param flagKey feature flag identifying key
	 * @param context lookup context conditions
	 * @return feature flag value
	 */
	FlagValueByKey(flagKey string, context map[string]interface{}) (*response.FeatureFlagVariationValue, error)

	/**
	 * Returns the boolean value of a feature flag for a given flag key, in an object that also describes the way the
	 * value was determined.
	 *
	 * @param flagKey the unique feature key for the feature flag.
	 * @param context the context for the feature flag lookup
	 * @param default the default value for if the flag value is unavailable.
	 * @param useDefaultContext whether to use the default context
	 * @return: an `FlagEvaluationDetail` object
	 */
	BooleanFlagVariationDetail(flagKey string, context map[string]interface{}, defaultValue bool, useDefaultContext bool) (*response.FlagEvaluationDetail[bool], error)

	/**
	 * Returns the long value of a feature flag for a given flag key, in an object that also describes the way the
	 * value was determined.
	 *
	 * @param flagKey the unique feature key for the feature flag.
	 * @param context the context for the feature flag lookup
	 * @param default the default value for if the flag value is unavailable.
	 * @param useDefaultContext whether to use the default context
	 * @return: an `FlagEvaluationDetail` object
	 */
	LongFlagVariationDetail(flagKey string, context map[string]interface{}, defaultValue int64, useDefaultContext bool) (*response.FlagEvaluationDetail[int64], error)

	/**
	 * Returns the string value of a feature flag for a given flag key, in an object that also describes the way the
	 * value was determined.
	 *
	 * @param flagKey the unique feature key for the feature flag.
	 * @param context the context for the feature flag lookup
	 * @param default the default value for if the flag value is unavailable.
	 * @param useDefaultContext whether to use the default context
	 * @return: an `FlagEvaluationDetail` object
	 */
	StringFlagVariationDetail(flagKey string, context map[string]interface{}, defaultValue string, useDefaultContext bool) (*response.FlagEvaluationDetail[string], error)

	/**
	 * Returns the double value of a feature flag for a given flag key, in an object that also describes the way the
	 * value was determined.
	 *
	 * @param flagKey the unique feature key for the feature flag.
	 * @param context the context for the feature flag lookup
	 * @param default the default value for if the flag value is unavailable.
	 * @param useDefaultContext whether to use the default context
	 * @return: an `FlagEvaluationDetail` object
	 */
	DoubleFlagVariationDetail(flagKey string, context map[string]interface{}, defaultValue float64, useDefaultContext bool) (*response.FlagEvaluationDetail[float64], error)
}

func NewNumeratorClient(config *config.NumeratorConfig) *DefaultNumeratorClient {
	httpClient := network.NewHttpClient(config)
	nService := service.NewNumeratorService(httpClient)
	return &DefaultNumeratorClient{
		service:         nService,
		contextProvider: context.NewContextProvider(),
	}
}
