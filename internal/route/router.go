package route

import (
	"context"

	"github.com/c0x12c/numerator-go-sdk/internal/clients"
	"github.com/c0x12c/numerator-go-sdk/internal/service"
	"github.com/c0x12c/numerator-go-sdk/pkg/config"
	"github.com/c0x12c/numerator-go-sdk/pkg/network"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Router interface {
	// Configure configures the router
	Configure(e *echo.Echo)
}

func Routers(ctx context.Context, mocks []any) (routes []Router,
	services []any,
	clts []any) {
	//// Repositories

	//// Client
	// Initialize Numerator configuration
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// apiKey := viper.Get("API_KEY").(string)
	apiKey := "NUM.oztqpZ2d7wmsAewW0sBKcQ==.wUd8cUDl4uytg3TmHmtl4sKzVrMkEfbvOMQGRP/xurNiuVOBWpsgDJuScQmSdKdi"
	numeratorConfig := config.NewNumeratorConfig(apiKey)
	httpClient := network.NewHttpClient(numeratorConfig)

	// Create Numerator client
	numeratorClient := clients.NewNumeratorClient(apiKey, numeratorConfig)

	//// Services
	numeratorService := service.NewNumeratorService(httpClient)

	return []Router{
			NewNumeratorRouter(numeratorService),
		}, []any{
			numeratorService,
		}, []any{
			numeratorClient,
		}
}
