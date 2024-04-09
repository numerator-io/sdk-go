package main

import (
	"fmt"

	"github.com/numerator-io/sdk-go/pkg/context"
	"github.com/numerator-io/sdk-go/pkg/log"

	"github.com/numerator-io/sdk-go/internal/clients"
	"github.com/numerator-io/sdk-go/pkg/config"
	"github.com/numerator-io/sdk-go/pkg/constant"
)

// NumeratorProvider embeds existed providers
type NumeratorProvider struct {
	*clients.NumeratorFeatureFlagProvider
}

// NewNumeratorProvider creates a new instance of NumeratorProvider
func NewNumeratorProvider(config *config.NumeratorConfig, contextProvider context.ContextProvider) *NumeratorProvider {
	return &NumeratorProvider{
		clients.NewNumeratorFeatureFlagProvider(config, contextProvider),
	}
}

// Custom method for dev use
func (p *NumeratorProvider) FeatureFlag2Enabled() string {
	// You may choose to fetch feature flag of type BOOLEAN, STRING, LONG (int64) or DOUBLE (float64)
	flagKey := "" // add your flag key here
	defaultValue := "default"
	givenContext := map[string]interface{}{}
	useDefaultContext := true
	result := p.GetStringFeatureFlag(flagKey, defaultValue, givenContext, useDefaultContext)
	return result
}

func main() {
	// Initialize Numerator configuration
	apiKey := "" // add your apiKey here
	numeratorConfig := config.NewNumeratorConfig(apiKey)

	// Create a log instance
	logger, _ := log.NewZapLogger()

	// Create Numerator client
	contextProvider := context.NewContextProvider()
	numeratorClient := clients.NewNumeratorClient(numeratorConfig, contextProvider)

	/**** EXAMPLE USING CLIENT ****/

	// Fetch feature flag list
	flags, err := numeratorClient.FeatureFlags(constant.Page, constant.Size)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to fetch feature flags: %v", err))
		return
	}
	logger.Info(fmt.Sprintf("Fetched Feature Flags: %v", flags))

	// Fetch feature flag details
	flagKey := "" // add your flag key here
	flagDetail, err := numeratorClient.FeatureFlagDetails(flagKey)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to fetch feature flag details: %v", err))
		return
	}
	logger.Info(fmt.Sprintf("Fetched Feature Flag Details: %v", flagDetail))

	// Fetch feature flag value by key with empty context
	// You may choose to fetch feature flag of type BOOLEAN, STRING, LONG (int64) or DOUBLE (float64)
	defaultValue := "default"
	gotValue, err := numeratorClient.StringFlagVariationDetail(flagKey, nil, defaultValue, false)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to fetch flag value by key: %v", err))
		return
	}
	logger.Info(fmt.Sprintf("Fetched Flag Value by Key: %v", gotValue))

	// Create a context
	contextEnv := map[string]interface{}{
		"env": "dev",
	}

	// Fetch feature flag value by key with context
	gotValue, err = numeratorClient.StringFlagVariationDetail(flagKey, contextEnv, defaultValue, false)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to fetch flag value by key with context: %v", err))
		return
	}
	logger.Info(fmt.Sprintf("Fetched Flag Value by Key with Context: %v", gotValue))

	/**** EXAMPLE USING FEATURE FLAG PROVIDER ****/

	contextProvider.Set("env", "dev")
	exampleNumeratorProvider := NewNumeratorProvider(numeratorConfig, contextProvider)
	gotStringValue := exampleNumeratorProvider.FeatureFlag2Enabled()
	logger.Info(fmt.Sprintf("Use provider to fetch flag value: %v", gotStringValue))
}
