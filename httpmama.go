// Package httpmama provides a set of utility to work with HTTP
package httpmama

import (
	"net/http"
	"net/http/httptest"
)

// TestEndpoint store the endpoint specific configuration
type TestEndpoint struct {
	Path           string
	ResponseString string
	ResponseHeader http.Header
}

// ServerConfig store server configuration
type ServerConfig struct {
	TestEndpoints []TestEndpoint
}

// NewTestServer create new HTTP server for testing
func NewTestServer(config ServerConfig) *httptest.Server {
	mux := http.NewServeMux()

	for _, endpoint := range config.TestEndpoints {
		path := endpoint.Path
		responseString := endpoint.ResponseString
		responseHeader := endpoint.ResponseHeader

		handler := func(w http.ResponseWriter, r *http.Request) {
			for key, values := range responseHeader {
				for _, value := range values {
					w.Header().Add(key, value)
				}
			}
			w.Write([]byte(responseString))
		}

		mux.HandleFunc(path, handler)
	}

	return httptest.NewServer(mux)
}
