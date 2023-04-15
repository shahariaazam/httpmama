package httpmama

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTestServer(t *testing.T) {
	testEndpoints := []TestEndpoint{
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
		{
			Path:           "/user",
			ResponseString: "hello, john!",
			ResponseHeader: http.Header{"Content-Type": []string{"text/plain"}},
			QueryParams:    url.Values{"name": []string{"john"}},
		},
		{
			Path:           "/user",
			ResponseString: "hello, doe!",
			ResponseHeader: http.Header{"Content-Type": []string{"text/plain"}},
			QueryParams:    url.Values{"name": []string{"doe"}},
		},
		{
			Path:           "/user",
			ResponseString: "hello, doe! Age: 30",
			ResponseHeader: http.Header{"Content-Type": []string{"text/plain"}},
			QueryParams:    url.Values{"name": []string{"doe"}, "age": []string{"30"}},
		},
		{
			Path:           "/return-json",
			ResponseString: `{"hello": "world"}`,
			ResponseHeader: http.Header{"Content-Type": []string{"application/json"}, "Some-Other-Header": []string{"some-other-header"}},
		},
	}

	testCases := []struct {
		name               string
		endpoints          []TestEndpoint
		requestPath        string
		expectedBody       string
		expectedHeaders    map[string]string
		expectedStatusCode int
	}{
		{
			name:            "single endpoint",
			requestPath:     "/foo",
			expectedBody:    "hello, world!",
			expectedHeaders: map[string]string{"Content-Type": "text/plain"},
		},
		{
			name:            "endpoint with single query params",
			requestPath:     "/user?" + url.Values{"name": []string{"doe"}}.Encode(),
			expectedBody:    "hello, doe!",
			expectedHeaders: map[string]string{"Content-Type": "text/plain"},
		},
		{
			name:            "endpoint with multiple query params",
			requestPath:     "/user?" + url.Values{"name": []string{"doe"}, "age": []string{"30"}}.Encode(),
			expectedBody:    "hello, doe! Age: 30",
			expectedHeaders: map[string]string{"Content-Type": "text/plain"},
		},
		{
			name:            "non-existent endpoint should return 404",
			requestPath:     "/something/else",
			expectedBody:    "404 page not found\n",
			expectedHeaders: map[string]string{"Content-Type": "text/plain; charset=utf-8"},
		},
		{
			name:            "returning json with correct header",
			requestPath:     "/return-json",
			expectedBody:    `{"hello": "world"}`,
			expectedHeaders: map[string]string{"Content-Type": "application/json"},
		},
		{
			name:            "return multiple header",
			requestPath:     "/return-json",
			expectedBody:    `{"hello": "world"}`,
			expectedHeaders: map[string]string{"Content-Type": "application/json", "Some-Other-Header": "some-other-header"},
		},
	}

	config := ServerConfig{TestEndpoints: testEndpoints}
	server := NewTestServer(config)
	defer server.Close()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(server.URL + tc.requestPath)
			if err != nil {
				t.Errorf("error making GET request: %v", err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("error reading response body: %v", err)
			}

			assert.Equal(t, tc.expectedBody, string(body))

			for hk, hv := range tc.expectedHeaders {
				assert.Equal(t, hv, resp.Header.Get(hk))
			}
		})
	}
}
