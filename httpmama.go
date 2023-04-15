// Package httpmama provides a set of utility to work with HTTP
package httpmama

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

// TestEndpoint store the endpoint specific configuration
type TestEndpoint struct {
	Path           string
	ResponseString string
	ResponseHeader http.Header
	QueryParams    url.Values
}

// ServerConfig store server configuration
type ServerConfig struct {
	TestEndpoints []TestEndpoint
}

// NewTestServer create new HTTP server for testing
func NewTestServer(config ServerConfig) *httptest.Server {
	mux := http.NewServeMux()

	// Use a map to store the response for each path and query parameter
	responses := map[string]map[string]TestEndpoint{}

	// Register a handler that looks up the response based on the path and query parameters
	handler := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		queryParams := r.URL.Query()

		response, ok := responses[path][queryParams.Encode()]
		if !ok {
			// No response found for this path and query parameters, return a 404 response
			http.NotFound(w, r)
			return
		}

		for key, values := range response.ResponseHeader {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		w.Write([]byte(response.ResponseString))
	}

	// Register the handler for all paths
	mux.HandleFunc("/", handler)

	// Register each endpoint with its own set of query parameters
	for _, endpoint := range config.TestEndpoints {
		path := endpoint.Path

		if responses[path] == nil {
			// Create a new map for this path if it doesn't exist yet
			responses[path] = map[string]TestEndpoint{}
		}

		// Store the response for this path and query parameters
		responses[path][endpoint.QueryParams.Encode()] = endpoint
	}

	return httptest.NewServer(mux)
}
