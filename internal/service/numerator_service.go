package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/c0x12c/numerator-go-sdk/internal/pkg/api/request"
	"github.com/c0x12c/numerator-go-sdk/internal/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/internal/pkg/exception"
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
	requestBody := request.FlagValueByKeyRequest{
		Key:     flagKey,
		Context: context,
	}
	resp, err := s.HttpClient.Post(FLAG_VALUE_BY_KEY, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	return s.handleResponse(resp)
}

func (s *NumeratorService) FlagList(page, size int) (response.ApiResponse, error) {
	requestBody := request.FlagListRequest{
		Page: page,
		Size: size,
	}
	resp, err := s.HttpClient.Post(FLAG_LISTING, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	return s.handleResponse(resp)
}

func (s *NumeratorService) FlagDetailByKey(flagKey string) (response.ApiResponse, error) {
	queryParams := map[string]string{"key": flagKey}
	resp, err := s.HttpClient.Post(FLAG_DETAIL_BY_KEY, queryParams, "")
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
		msg := exception.NumeratorLogMessage.INVALID_SDK_KEY_ERROR
		numeratorError.Message = &msg
	case 404:
		msg := exception.GetObjectDoesNotExist(*numeratorError.Message)
		numeratorError.Message = &msg
	case 400:
		msg := exception.NumeratorLogMessage.BAD_REQUEST_ERROR
		numeratorError.Message = &msg
	default:
		msg := exception.GetUnexpectedHttpResponse(*numeratorError.Message)
		numeratorError.Message = &msg
	}

	return &response.FailureResponse{Error: numeratorError}, nil
}
