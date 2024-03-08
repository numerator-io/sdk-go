package main

import (
	"log"

	"github.com/c0x12c/numerator-go-sdk/internal/clients"
	"github.com/c0x12c/numerator-go-sdk/pkg/config"
	"github.com/c0x12c/numerator-go-sdk/pkg/constant"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// Initialize Numerator configuration
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	apiKey := viper.Get("API_KEY").(string)
	numeratorConfig := config.NewNumeratorConfig(apiKey)

	// Create a logger instance
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	defer logger.Sync()

	// Create Numerator client
	numeratorClient := clients.NewDefaultNumeratorClient(numeratorConfig)

	// Fetch feature flags
	flags, err := numeratorClient.FeatureFlags(constant.Page, constant.Size)
	if err != nil {
		logger.Error("failed to fetch feature flags", zap.Error(err))
		return
	}
	logger.Info("fetched feature flags", zap.Any("flags", flags))

	// Fetch feature flag details
	flagKey := "go_featureflag_01"
	flagDetail, err := numeratorClient.FeatureFlagDetails(flagKey)
	if err != nil {
		logger.Error("failed to fetch feature flag details", zap.Error(err))
		return
	}
	logger.Info("fetched feature flag details", zap.Any("flagDetail", flagDetail))

	// Create an empty context
	context := make(map[string]interface{})

	// Fetch feature flag value by key with empty context
	defaultBoolean := true
	booleanValue, err := numeratorClient.GetValueByKeyWithDefault(flagKey, context, defaultBoolean)
	if err != nil {
		logger.Error("failed to fetch flag value by key", zap.Error(err))
		return
	}
	// Perform type assertion to ensure booleanValue is a boolean
	boolVal, ok := booleanValue.(bool)
	if !ok {
		logger.Error("failed to fetch flag value by key: booleanValue is not a boolean")
		return
	}
	logger.Info("fetched flag value by key", zap.Bool("BooleanValue", boolVal))

	// Create a context
	contextEnv := map[string]interface{}{
		"env": "dev",
	}

	// Fetch feature flag value by key with empty context
	booleanValue, err = numeratorClient.GetValueByKeyWithDefault(flagKey, contextEnv, defaultBoolean)
	if err != nil {
		logger.Error("failed to fetch flag value by key", zap.Error(err))
		return
	}
	// Perform type assertion to ensure booleanValue is a boolean
	boolVal, ok = booleanValue.(bool)
	if !ok {
		logger.Error("failed to fetch flag value by key: booleanValue is not a boolean")
		return
	}
	logger.Info("fetched flag value by key with context", zap.Bool("BooleanValue", boolVal))
}
