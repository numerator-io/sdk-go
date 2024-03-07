package main

import (
	"fmt"

	"github.com/c0x12c/numerator-go-sdk/internal/clients"
	"github.com/c0x12c/numerator-go-sdk/pkg/config"
)

func main() {
	apiKey := "NUM.oztqpZ2d7wmsAewW0sBKcQ==.wUd8cUDl4uytg3TmHmtl4sKzVrMkEfbvOMQGRP/xurNiuVOBWpsgDJuScQmSdKdi"
	// Initialize Numerator configuration
	numeratorConfig := config.NewNumeratorConfig(apiKey)

	// Create Numerator client
	numeratorClient := clients.NewDefaultNumeratorClient(numeratorConfig)

	// Fetch feature flags
	page := 0
	size := 10
	flags, err := numeratorClient.FeatureFlags(page, size)
	if err != nil {
		fmt.Println("Error fetching feature flags:", err)
		return
	}
	fmt.Println("Feature Flags:", flags)

	// Fetch feature flag details
	flagKey := "go_featureflag_01"
	flagDetail, err := numeratorClient.FeatureFlagDetails(flagKey)
	if err != nil {
		fmt.Println("Error fetching feature flag details:", err)
		return
	}
	fmt.Println("Feature Flag Detail:", *flagDetail)

	// Create an empty context
	context := make(map[string]interface{})

	// Fetch feature flag value by key with empty context
	defaultBoolean := true
	booleanValue, err := numeratorClient.GetValueByKeyWithDefault(flagKey, context, defaultBoolean)
	if err != nil {
		fmt.Println("Error fetching flag value by key:", err)
		return
	}
	fmt.Println("Boolean Value:", booleanValue)

	// Create a context
	contextEnv := map[string]interface{}{
		"env": "dev",
	}

	// Fetch feature flag value by key with empty context
	booleanValue, err = numeratorClient.GetValueByKeyWithDefault(flagKey, contextEnv, defaultBoolean)
	if err != nil {
		fmt.Println("Error fetching flag value by key:", err)
		return
	}
	fmt.Println("Boolean Value:", booleanValue)
}
