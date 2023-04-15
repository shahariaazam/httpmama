package httpmama

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestCreateTestServer(t *testing.T) {
	testCases := []struct {
		name            string
		endpoints       []TestEndpoint
		requestPath     string
		expectedBody    string
		expectedHeaders http.Header
	}{
		{
			name: "single endpoint",
			endpoints: []TestEndpoint{
				{
					Path:           "/foo",
					ResponseString: "hello, world!",
					ResponseHeader: http.Header{"Content-Type": []string{"text/plain"}},
				},
			},
			requestPath:     "/foo",
			expectedBody:    "hello, world!",
			expectedHeaders: http.Header{"Content-Type": []string{"text/plain"}},
		},
		{
			name: "multiple endpoints",
			endpoints: []TestEndpoint{
				{
					Path:           "/foo",
					ResponseString: "hello, world!",
					ResponseHeader: http.Header{"Content-Type": []string{"text/plain"}},
				},
				{
					Path:           "/bar",
					ResponseString: "goodbye, world!",
					ResponseHeader: http.Header{"Content-Type": []string{"text/plain"}},
				},
			},
			requestPath:     "/bar",
			expectedBody:    "goodbye, world!",
			expectedHeaders: http.Header{"Content-Type": []string{"text/plain"}},
		},
		{
			name: "endpoint with query params",
			endpoints: []TestEndpoint{
				{
					Path:           "/user",
					ResponseString: "hello, john!",
					ResponseHeader: http.Header{"Content-Type": []string{"text/plain"}},
					QueryParams:    url.Values{"name": []string{"john"}},
				},
			},
			requestPath:     "/user?" + url.Values{"name": []string{"john"}}.Encode(),
			expectedBody:    "hello, john!",
			expectedHeaders: http.Header{"Content-Type": []string{"text/plain"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config := ServerConfig{TestEndpoints: tc.endpoints}
			server := NewTestServer(config)
			defer server.Close()

			resp, err := http.Get(server.URL + tc.requestPath)
			if err != nil {
				t.Errorf("error making GET request: %v", err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("error reading response body: %v", err)
			}

			if string(body) != tc.expectedBody {
				t.Errorf("expected response body to be '%s', got '%s'", tc.expectedBody, string(body))
			}

			if resp.Header.Get("Content-Type") != "text/plain" {
				t.Errorf("expected Content-Type header to be 'text/plain', got '%s'", resp.Header.Get("Content-Type"))
			}
		})
	}
}
