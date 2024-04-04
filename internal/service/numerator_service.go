//go:generate mockgen -source numerator_service.go -destination ./mock_service/mock_numerator_service/numerator_service_gen.go
package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/c0x12c/numerator-go-sdk/pkg/api/request"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/pkg/exception"
	"github.com/c0x12c/numerator-go-sdk/pkg/network"
)

type NumeratorService interface {
	FlagValueByKey(requestBody request.FlagValueByKeyRequest) (response.ApiResponse, error)
	FlagList(requestBody request.FlagListRequest) (response.ApiResponse, error)
	FlagDetailByKey(flagKey string) (response.ApiResponse, error)
}

type DefaultNumeratorService struct {
	HttpClient network.HttpClient // Use the HttpClient from the network package
}

func NewNumeratorService(httpClient network.HttpClient) NumeratorService {
	return &DefaultNumeratorService{
		HttpClient: httpClient,
	}
}

func (s *DefaultNumeratorService) FlagValueByKey(requestBody request.FlagValueByKeyRequest) (response.ApiResponse, error) {
	resp, err := s.HttpClient.Post(FLAG_VALUE_BY_KEY, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	return s.handleResponse(resp, response.FeatureFlagVariationValue{})
}

func (s *DefaultNumeratorService) FlagList(requestBody request.FlagListRequest) (response.ApiResponse, error) {
	resp, err := s.HttpClient.Post(FLAG_LISTING, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	return s.handleResponse(resp, response.FeatureFlagList{})
}

func (s *DefaultNumeratorService) FlagDetailByKey(flagKey string) (response.ApiResponse, error) {
	queryParams := map[string]string{"key": flagKey}
	resp, err := s.HttpClient.Post(FLAG_DETAIL_BY_KEY, queryParams, "")
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	return s.handleResponse(resp, response.FeatureFlag{})
}

func (s *DefaultNumeratorService) handleResponse(resp *http.Response, respType interface{}) (response.ApiResponse, error) {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		switch respType.(type) {
		case response.FeatureFlagList:
			featureFlag := new(response.FeatureFlagList)
			err := json.NewDecoder(resp.Body).Decode(&featureFlag)
			if err != nil {
				return nil, fmt.Errorf("failed to decode JSON response: %v", err)
			}
			return &response.SuccessResponse[response.FeatureFlagList]{SuccessResponse: *featureFlag}, nil
		case response.FeatureFlag:
			featureFlag := new(response.FeatureFlag)
			err := json.NewDecoder(resp.Body).Decode(&featureFlag)
			if err != nil {
				return nil, fmt.Errorf("failed to decode JSON response: %v", err)
			}
			return &response.SuccessResponse[response.FeatureFlag]{SuccessResponse: *featureFlag}, nil
		case response.FeatureFlagVariationValue:
			featureFlag := new(response.FeatureFlagVariationValue)
			err := json.NewDecoder(resp.Body).Decode(&featureFlag)
			if err != nil {
				return nil, fmt.Errorf("failed to decode JSON response: %v", err)
			}
			return &response.SuccessResponse[response.FeatureFlagVariationValue]{SuccessResponse: *featureFlag}, nil
		default:
			return nil, fmt.Errorf("request type %s hasn't been supported yet", respType)
		}

	}

	var numeratorError response.NumeratorError
	err := json.NewDecoder(resp.Body).Decode(&numeratorError)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON error response: %v", err)
	}

	// Switch based on HTTP status code constants
	switch resp.StatusCode {
	case http.StatusUnauthorized:
		numeratorError.Message = exception.INVALID_SDK_KEY_ERROR
	case http.StatusNotFound:
		numeratorError.Message = exception.GetObjectDoesNotExist(numeratorError.Message)
	case http.StatusBadRequest:
		numeratorError.Message = exception.BAD_REQUEST_ERROR
	default:
		numeratorError.Message = exception.GetUnexpectedHttpResponse(numeratorError.Message)
	}

	return &response.FailureResponse{Error: numeratorError}, nil
}
