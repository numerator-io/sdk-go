# numerator-go-sdk
Numerator SDK for the Go programming language.

## Quickstart

1. Create a `NumeratorProvider` embedding prefered Providers and define some custom methods:

```go
import (
	"github.com/numerator-io/sdk-go/internal/clients"
	"github.com/numerator-io/sdk-go/pkg/context"
	"github.com/numerator-io/sdk-go/pkg/log"
	"github.com/numerator-io/sdk-go/pkg/config"
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
	flagKey := "featureflag_2" // add your flag key here
	defaultValue := "default"
	givenContext := map[string]interface{}{}
	useDefaultContext := true
	result := p.GetStringFeatureFlag(flagKey, defaultValue, givenContext, useDefaultContext)
	return result
}
```

2. Use the above feature flag accessors in code:

```golang
exampleNumeratorProvider := NewNumeratorProvider(numeratorConfig, contextProvider)
enabled := exampleNumeratorProvider.FeatureFlag2Enabled()

if enabled {
    // feature enabled logic
} else {
    // feature disabled logic
}
```

## Advanced Usage

*TODO*