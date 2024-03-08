package config

type NumeratorConfig struct {
	BaseURL        string
	APIKey         string
	ConnectTimeout int64 // In milliseconds
	ReadTimeout    int64 // In milliseconds
	WriteTimeout   int64 // In milliseconds
}

// NewNumeratorConfig creates a new NumeratorConfig with default values
func NewNumeratorConfig(apiKey string) *NumeratorConfig {
	return &NumeratorConfig{
		BaseURL:        "https://service-platform.dev.numerator.io",
		APIKey:         apiKey,
		ConnectTimeout: 10000, // default to 10 seconds
		ReadTimeout:    10000, // default to 10 seconds
		WriteTimeout:   10000, // default to 10 seconds
	}
}
