package main

import (
	"fmt"

	"github.com/c0x12c/numerator-go-sdk/internal/clients"
	"github.com/c0x12c/numerator-go-sdk/internal/pkg/config"
)

func main() {
	// Initialize Numerator configuration
	numeratorConfig := &config.NumeratorConfig{
		BaseURL:        "https://your-api-base-url.com",
		ConnectTimeout: 5000, // Adjust timeout values as needed
		ReadTimeout:    5000,
		WriteTimeout:   5000,
		APIKey:         "Your API Key",
	}

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
	flagKey := "test_string"
	flagDetail, err := numeratorClient.FeatureFlagDetails(flagKey)
	if err != nil {
		fmt.Println("Error fetching feature flag details:", err)
		return
	}
	fmt.Println("Feature Flag Detail:", flagDetail)

	// Fetch feature flag value by key with default values
	defaultBoolean := true
	booleanValue, err := numeratorClient.GetValueByKeyWithDefault(flagKey, nil, defaultBoolean)
	if err != nil {
		fmt.Println("Error fetching flag value by key:", err)
		return
	}
	fmt.Println("Boolean Value:", booleanValue)

	defaultLong := int64(10)
	longValue, err := numeratorClient.GetValueByKeyWithDefault("test_long", nil, defaultLong)
	if err != nil {
		fmt.Println("Error fetching flag value by key:", err)
		return
	}
	fmt.Println("Long Value:", longValue)

	defaultString := "default"
	stringValue, err := numeratorClient.GetValueByKeyWithDefault(flagKey, nil, defaultString)
	if err != nil {
		fmt.Println("Error fetching flag value by key:", err)
		return
	}
	fmt.Println("String Value:", stringValue)

	defaultDouble := 0.0
	doubleValue, err := numeratorClient.GetValueByKeyWithDefault(flagKey, nil, defaultDouble)
	if err != nil {
		fmt.Println("Error fetching flag value by key:", err)
		return
	}
	fmt.Println("Double Value:", doubleValue)
}
