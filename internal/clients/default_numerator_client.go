package clients

import (
	"fmt"

	"github.com/c0x12c/numerator-go-sdk/internal/service"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/request"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/pkg/config"
	"github.com/c0x12c/numerator-go-sdk/pkg/enum"
	"github.com/c0x12c/numerator-go-sdk/pkg/exception"
	"github.com/c0x12c/numerator-go-sdk/pkg/network"
)

type DefaultNumeratorClient struct {
	service service.NumeratorService
}

func NewDefaultNumeratorClient(config *config.NumeratorConfig) *DefaultNumeratorClient {
	httpClient := network.NewHttpClient(config)
	nService := service.NewNumeratorService(httpClient)
	return &DefaultNumeratorClient{
		service: nService,
	}
}

func (c *DefaultNumeratorClient) FeatureFlags(page, size int) ([]response.FeatureFlag, error) {
	requestBody := request.FlagListRequest{
		Page: page,
		Size: size,
	}
	resp, err := c.service.FlagList(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch feature flags: %v", err)
	}
	successResp, ok := resp.(*response.SuccessResponse[response.FeatureFlagListResponse])
	if !ok {
		return nil, handleErrorResponse(resp)
	}
	flagListResp := successResp.SuccessResponse
	return flagListResp.Data(), nil
}

func (c *DefaultNumeratorClient) FeatureFlagDetails(flagKey string) (*response.FeatureFlag, error) {
	resp, err := c.service.FlagDetailByKey(flagKey)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch feature flag details: %v", err)
	}
	successResp, ok := resp.(*response.SuccessResponse[response.FeatureFlag])
	if !ok {
		return nil, handleErrorResponse(resp)
	}
	flagDetailResp := &successResp.SuccessResponse
	return flagDetailResp, nil
}

func (c *DefaultNumeratorClient) GetValueByKeyWithDefault(flagKey string, context map[string]interface{}, defaultValue interface{}) (interface{}, error) {
	requestBody := request.FlagValueByKeyRequest{
		Key:     flagKey,
		Context: context,
	}
	resp, err := c.service.FlagValueByKey(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch flag value by key: %v", err)
	}
	successResp, ok := resp.(*response.SuccessResponse[response.FeatureFlagVariationValue])
	if !ok {
		return nil, handleErrorResponse(resp)
	}
	flagValueResp := &successResp.SuccessResponse
	return convertVariationValue(flagValueResp, defaultValue)
}

func convertVariationValue(flagVariationValue *response.FeatureFlagVariationValue, defaultValue interface{}) (interface{}, error) {
	switch defaultValue.(type) {
	case string:
		if flagVariationValue.ValueType != enum.STRING {
			return nil, fmt.Errorf("type mismatch")
		}
		return flagVariationValue.Value.StringValue, nil
	case bool:
		if flagVariationValue.ValueType != enum.BOOLEAN {
			return nil, fmt.Errorf("type mismatch")
		}
		return flagVariationValue.Value.BooleanValue, nil
	case int:
		if flagVariationValue.ValueType != enum.LONG {
			return nil, fmt.Errorf("type mismatch")
		}
		return flagVariationValue.Value.LongValue, nil
	case float64:
		if flagVariationValue.ValueType != enum.DOUBLE {
			return nil, fmt.Errorf("type mismatch")
		}
		return flagVariationValue.Value.DoubleValue, nil
	default:
		return nil, fmt.Errorf("unsupported type")
	}
}

func handleErrorResponse(resp response.ApiResponse) error {
	failureResp, ok := resp.(*response.FailureResponse)
	if !ok {
		return fmt.Errorf("unexpected response format")
	}
	return &exception.NumeratorException{Message: failureResp.Error.Message, Status: failureResp.Error.HttpStatus}
}
