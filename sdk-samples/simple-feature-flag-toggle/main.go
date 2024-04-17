package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/numerator-io/sdk-go/internal/clients"
	"github.com/numerator-io/sdk-go/pkg/config"
	"github.com/numerator-io/sdk-go/pkg/context"
)

type NumeratorProvider struct {
	*clients.NumeratorFeatureFlagProvider
}

func NewNumeratorProvider(config *config.NumeratorConfig, contextProvider context.ContextProvider) *NumeratorProvider {
	return &NumeratorProvider{
		clients.NewNumeratorFeatureFlagProvider(config, contextProvider),
	}
}

func (p *NumeratorProvider) FeatureFlagHelloEnabled() bool {
	flagKey := "hello_logic"
	defaultValue := false
	givenContext := map[string]interface{}{} // empty context
	useDefaultContext := false
	enabled := p.GetBooleanFeatureFlag(flagKey, defaultValue, givenContext, useDefaultContext)
	return enabled
}

func (p *NumeratorProvider) GetUserCountry(environment string) string {
	flagKey := "user_country"
	defaultValue := "vn"
	givenContext := map[string]interface{}{"environment": environment}
	useDefaultContext := false
	country := p.GetStringFeatureFlag(flagKey, defaultValue, givenContext, useDefaultContext)
	return country
}

func (p *NumeratorProvider) GetTokenExpiration(userId string, userEmail string) int64 {
	flagKey := "expiration_token"
	defaultValue := int64(0)
	givenContext := map[string]interface{}{"user_id": userId, "user_email": userEmail}
	useDefaultContext := false
	token_expiration := p.GetLongFeatureFlag(flagKey, defaultValue, givenContext, useDefaultContext)
	return token_expiration
}

func main() {
	// Update your apiKey here
	apiKey := "NUM.6CCARNz+Klhj2My9BYM9tA==.19sxPRuTQQ7pyUoZ/qvnHyFXWnAL7b68GWHfB8F2Pz/pFSXty7XFNb209DwUHifp"
	numeratorConfig := config.NewNumeratorConfig(apiKey)
	contextProvider := context.NewContextProvider()
	// contextProvider.Set("env", "dev")
	exampleNumeratorProvider := NewNumeratorProvider(numeratorConfig, contextProvider)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", func(c echo.Context) error {
		enabled := exampleNumeratorProvider.FeatureFlagHelloEnabled()
		if enabled {
			return c.JSON(http.StatusOK, fmt.Sprintf("Flag value is %t. Welcome to Numerator!", enabled))
		}
		return c.JSON(http.StatusOK, fmt.Sprintf("Flag value is %t.", enabled))
	})

	e.GET("/user-country", func(c echo.Context) error {
		environment := c.QueryParam("environment")
		country := exampleNumeratorProvider.GetUserCountry(environment)
		return c.JSON(http.StatusOK, fmt.Sprintf("Country is %s.", country))
	})

	e.GET("/expiration-token", func(c echo.Context) error {
		id := c.QueryParam("userId")
		email := c.QueryParam("userEmail")
		tokenExpiration := exampleNumeratorProvider.GetTokenExpiration(id, email)
		return c.JSON(http.StatusOK, fmt.Sprintf("Expiration token is %v seconds.", tokenExpiration))
	})

	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	select {}
}
