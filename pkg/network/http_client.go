//go:generate mockgen -source http_client.go -destination ./mock_http_client/http_client_gen.go
package network

import "net/http"

type HttpClient interface {
	Post(path string, queryParams map[string]string, body interface{}) (*http.Response, error)
}
