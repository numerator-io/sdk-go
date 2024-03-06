package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/c0x12c/numerator-go-sdk/internal/pkg/api/request"
	"github.com/c0x12c/numerator-go-sdk/internal/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/internal/pkg/network"
)

type NumeratorService struct {
	HttpClient *network.HttpClient // Use the HttpClient from the network package
}

func NewNumeratorService(httpClient *network.HttpClient) *NumeratorService {
	return &NumeratorService{
		HttpClient: httpClient,
	}
}

func (s *NumeratorService) FlagValueByKey(flagKey string, context map[string]interface{}) (response.ApiResponse, error) {
	requestBody := request.FlagByKeyRequest{
		Key:     flagKey,
		Context: context,
	}
	resp, err := s.HttpClient.Post(FlagValueByKey, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	return s.handleResponse(resp)
}

func (s *NumeratorService) handleResponse(resp *http.Response) (response.ApiResponse, error) {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		var featureFlag response.FeatureFlag
		err := json.NewDecoder(resp.Body).Decode(&featureFlag)
		if err != nil {
			return nil, fmt.Errorf("failed to decode JSON response: %v", err)
		}
		return &response.SuccessResponse{SuccessResponse: featureFlag}, nil
	}

	var numeratorError response.NumeratorError
	err := json.NewDecoder(resp.Body).Decode(&numeratorError)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON error response: %v", err)
	}

	switch resp.StatusCode {
	case 401:
		msg := "API key is invalid"
		numeratorError.Message = &msg // Convert string constant to pointer to string
	default:
		if numeratorError.Message == nil {
			status := resp.Status
			numeratorError.Message = &status // Convert response status to pointer to string
		}
	}

	return &response.FailureResponse{Error: numeratorError}, nil
}
