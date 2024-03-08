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

type HttpClient struct {
	config  *config.NumeratorConfig
	client  *http.Client
	baseURL string
}

func NewHttpClient(config *config.NumeratorConfig) *HttpClient {
	client := &http.Client{
		Timeout: time.Duration(config.ConnectTimeout) * time.Millisecond,
	}
	return &HttpClient{
		config:  config,
		client:  client,
		baseURL: config.BaseURL,
	}
}

func (c *HttpClient) Post(path string, queryParams map[string]string, body interface{}) (*http.Response, error) {
	url := c.buildUrl(path, queryParams)

	// Marshal the body to JSON
	jsonBody, err := json.Marshal(body)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body to JSON: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}
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

func (c *HttpClient) buildUrl(path string, queryParams map[string]string) string {
	url := c.baseURL + path

	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		// Remove the last "&" character
		url = url[:len(url)-1]
	}

	return url
}
