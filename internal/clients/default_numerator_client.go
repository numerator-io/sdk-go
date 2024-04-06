package clients

import (
	"fmt"
	"net/http"

	"github.com/c0x12c/numerator-go-sdk/internal/models"
	"github.com/c0x12c/numerator-go-sdk/internal/service"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/request"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/pkg/context"
	"github.com/c0x12c/numerator-go-sdk/pkg/enum"
	"github.com/c0x12c/numerator-go-sdk/pkg/exception"
)

type DefaultNumeratorClient struct {
	service         service.NumeratorService
	contextProvider context.ContextProvider
}

func (c *DefaultNumeratorClient) GetContextProvider() context.ContextProvider {
	return c.contextProvider
}

func (c *DefaultNumeratorClient) SetContextProvider(contextProvider context.ContextProvider) {
	c.contextProvider = contextProvider
}

func (c *DefaultNumeratorClient) Version() string {
	return "1.0.0" // Implement version retrieval here if needed
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
	successResp, ok := resp.(*response.SuccessResponse[response.FeatureFlagList])
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

func (c *DefaultNumeratorClient) FlagValueByKey(flagKey string, context map[string]interface{}) (*response.FeatureFlagVariationValue, error) {
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
	return flagValueResp, nil
}

func (c *DefaultNumeratorClient) BooleanFlagVariationDetail(flagKey string, context map[string]interface{}, defaultValue bool, useDefaultContext bool) (*response.FlagEvaluationDetail[bool], error) {
	requestContext := context
	if useDefaultContext {
		requestContext = c.contextProvider.Context()
	}
	resp, err := c.FlagValueByKey(flagKey, requestContext)
	if err != nil {
		return nil, err
	}
	result, ok := convertVariationValue(resp, defaultValue).(*response.FlagEvaluationDetail[bool])
	if !ok {
		return nil, handleError("flag value type conversion failed")
	}
	return result, nil
}

func (c *DefaultNumeratorClient) LongFlagVariationDetail(flagKey string, context map[string]interface{}, defaultValue int64, useDefaultContext bool) (*response.FlagEvaluationDetail[int64], error) {
	requestContext := context
	if useDefaultContext {
		requestContext = c.contextProvider.Context()
	}
	resp, err := c.FlagValueByKey(flagKey, requestContext)
	if err != nil {
		return nil, err
	}
	result, ok := convertVariationValue(resp, defaultValue).(*response.FlagEvaluationDetail[int64])
	if !ok {
		return nil, handleError("flag value type conversion failed")
	}
	return result, nil
}

func (c *DefaultNumeratorClient) StringFlagVariationDetail(flagKey string, context map[string]interface{}, defaultValue string, useDefaultContext bool) (*response.FlagEvaluationDetail[string], error) {
	requestContext := context
	if useDefaultContext {
		requestContext = c.contextProvider.Context()
	}
	resp, err := c.FlagValueByKey(flagKey, requestContext)
	if err != nil {
		return nil, err
	}
	result, ok := convertVariationValue(resp, defaultValue).(*response.FlagEvaluationDetail[string])
	if !ok {
		return nil, handleError("flag value type conversion failed")
	}
	return result, nil
}

func (c *DefaultNumeratorClient) DoubleFlagVariationDetail(flagKey string, context map[string]interface{}, defaultValue float64, useDefaultContext bool) (*response.FlagEvaluationDetail[float64], error) {
	requestContext := context
	if useDefaultContext {
		requestContext = c.contextProvider.Context()
	}
	resp, err := c.FlagValueByKey(flagKey, requestContext)
	if err != nil {
		return nil, err
	}
	result, ok := convertVariationValue(resp, defaultValue).(*response.FlagEvaluationDetail[float64])
	if !ok {
		return nil, handleError("flag value type conversion failed")
	}
	return result, nil
}

func convertVariationValue(flagVariationValue *response.FeatureFlagVariationValue, defaultValue interface{}) interface{} {
	switch value := defaultValue.(type) {
	case string:
		defaultString := value
		if flagVariationValue.ValueType != enum.STRING {
			return failedConversion(flagVariationValue.Key, defaultString, nil)
		}
		return &response.FlagEvaluationDetail[string]{
			Key:    flagVariationValue.Key,
			Value:  flagVariationValue.Value.StringValue,
			Reason: nil,
		}
	case bool:
		defaultBoolean := value
		if flagVariationValue.ValueType != enum.BOOLEAN {
			return failedConversion(flagVariationValue.Key, defaultBoolean, nil)
		}
		return &response.FlagEvaluationDetail[bool]{
			Key:    flagVariationValue.Key,
			Value:  flagVariationValue.Value.BooleanValue,
			Reason: nil,
		}
	case int64:
		defaultLong := value
		if flagVariationValue.ValueType != enum.LONG {
			return failedConversion(flagVariationValue.Key, defaultLong, nil)
		}
		return &response.FlagEvaluationDetail[int64]{
			Key:    flagVariationValue.Key,
			Value:  flagVariationValue.Value.LongValue,
			Reason: nil,
		}
	case float64:
		defaultDouble := value
		if flagVariationValue.ValueType != enum.DOUBLE {
			return failedConversion(flagVariationValue.Key, defaultDouble, nil)
		}
		return &response.FlagEvaluationDetail[float64]{
			Key:    flagVariationValue.Key,
			Value:  flagVariationValue.Value.DoubleValue,
			Reason: nil,
		}
	default:
		reason := map[string]interface{}{
			"kind":      "Error",
			"errorKind": "unsupported type",
		}
		return failedConversion(flagVariationValue.Key, models.Unknown{Value: defaultValue}, reason)
	}
}

func failedConversion[T models.FlagValueType](flagKey string, value T, reason map[string]interface{}) *response.FlagEvaluationDetail[T] {
	// default reason is type mismatch
	if reason == nil {
		reason = map[string]interface{}{
			"kind":      "Error",
			"errorKind": "type mismatch",
		}
	}
	return &response.FlagEvaluationDetail[T]{
		Key:    flagKey,
		Value:  value,
		Reason: reason,
	}
}

func handleErrorResponse(resp response.ApiResponse) error {
	failureResp, ok := resp.(*response.FailureResponse)
	if !ok {
		return handleError("unexpected response format")
	}
	return exception.NewNumeratorException(failureResp.Error.Message, failureResp.Error.HttpStatus)
}

func handleError(error string) error {
	return exception.NewNumeratorException(error, http.StatusInternalServerError)
}
