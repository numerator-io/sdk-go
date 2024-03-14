package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/c0x12c/numerator-go-sdk/pkg/config"
	"github.com/c0x12c/numerator-go-sdk/pkg/constant"
)

type DefaultHttpClient struct {
	config  *config.NumeratorConfig
	client  *http.Client
	baseURL string
}

var _ HttpClient = (*DefaultHttpClient)(nil)

func NewHttpClient(config *config.NumeratorConfig) HttpClient {
	client := &http.Client{
		Timeout: time.Duration(config.ConnectTimeout) * time.Millisecond,
	}
	return &DefaultHttpClient{
		config:  config,
		client:  client,
		baseURL: config.BaseURL,
	}
}

func (c *DefaultHttpClient) Post(path string, queryParams map[string]string, body interface{}) (*http.Response, error) {
	// Marshal the body to JSON
	jsonBody, err := json.Marshal(body)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body to JSON: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, c.buildUrl(path), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// Build query params
	q := req.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// Build headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Add(constant.XNumAPIKeyHeader, c.config.APIKey)

	// Perform the HTTP request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %v", err)
	}

	return resp, nil
}

func (c *DefaultHttpClient) buildUrl(path string) string {
	url := c.baseURL + path
	return url
}
