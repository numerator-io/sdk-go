package network

import "net/http"

type HttpClient interface {
	Post(path string, queryParams map[string]string, body interface{}) (*http.Response, error)
}
