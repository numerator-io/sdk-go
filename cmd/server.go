package main

import (
	"fmt"

	"github.com/c0x12c/numerator-go-sdk/internal/clients"
	"github.com/c0x12c/numerator-go-sdk/internal/pkg/config"
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

	// Create an empty context map
	context := make(map[string]interface{})

	// Fetch feature flag value by key with default values
	defaultBoolean := true
	booleanValue, err := numeratorClient.GetValueByKeyWithDefault(flagKey, context, defaultBoolean)
	if err != nil {
		fmt.Println("Error fetching flag value by key:", err)
		return
	}
	fmt.Println("Boolean Value:", booleanValue)

	defaultLong := int64(10)
	longValue, err := numeratorClient.GetValueByKeyWithDefault("test_long", context, defaultLong)
	if err != nil {
		fmt.Println("Error fetching flag value by key:", err)
		return
	}
	fmt.Println("Long Value:", longValue)

	defaultString := "default"
	stringValue, err := numeratorClient.GetValueByKeyWithDefault(flagKey, context, defaultString)
	if err != nil {
		fmt.Println("Error fetching flag value by key:", err)
		return
	}
	fmt.Println("String Value:", stringValue)

	defaultDouble := 0.0
	doubleValue, err := numeratorClient.GetValueByKeyWithDefault(flagKey, context, defaultDouble)
	if err != nil {
		fmt.Println("Error fetching flag value by key:", err)
		return
	}
	fmt.Println("Double Value:", doubleValue)
}
