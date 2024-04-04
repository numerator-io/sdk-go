package main

import (
	"github.com/c0x12c/numerator-go-sdk/pkg/context"
	"github.com/c0x12c/numerator-go-sdk/pkg/log"

	"github.com/c0x12c/numerator-go-sdk/internal/clients"
	"github.com/c0x12c/numerator-go-sdk/pkg/config"
	"github.com/c0x12c/numerator-go-sdk/pkg/constant"

	"github.com/spf13/viper"
)

func main() {
	// Initialize Numerator configuration
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	apiKey := viper.Get("API_KEY").(string)
	numeratorConfig := config.NewNumeratorConfig(apiKey)

	// Create a log instance
	logger, _ := log.NewZapLogger()

	// Create Numerator client
	contextProvider := context.NewContextProvider()
	numeratorClient := clients.NewNumeratorClient(numeratorConfig, contextProvider)

	// Fetch feature flags
	flags, err := numeratorClient.FeatureFlags(constant.Page, constant.Size)
	if err != nil {
		logger.Error("failed to fetch feature flags", log.Error(err))
		return
	}
	logger.Info("fetched feature flags", log.Any("flags", flags))

	// Fetch feature flag details
	flagKey := "go_featureflag_01"
	flagDetail, err := numeratorClient.FeatureFlagDetails(flagKey)
	if err != nil {
		logger.Error("failed to fetch feature flag details", log.Error(err))
		return
	}
	logger.Info("fetched feature flag details", log.Any("flagDetail", flagDetail))

	// Fetch feature flag value by key with empty context
	flagKey = "go_featureflag_02"
	defaultBoolean := true
	booleanValue, err := numeratorClient.BooleanFlagVariationDetail(flagKey, nil, defaultBoolean, true)
	if err != nil {
		logger.Error("failed to fetch flag value by key", log.Error(err))
		return
	}
	logger.Info("fetched flag value by key", log.Any("BooleanValue", booleanValue))

	// Create a context
	contextEnv := map[string]interface{}{
		"env": "dev",
	}

	// Fetch feature flag value by key with empty context
	booleanValue, err = numeratorClient.BooleanFlagVariationDetail(flagKey, contextEnv, defaultBoolean, false)
	if err != nil {
		logger.Error("failed to fetch flag value by key", log.Error(err))
		return
	}
	logger.Info("fetched flag value by key with context", log.Any("BooleanValue", booleanValue))
}
