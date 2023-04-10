package httpmama

import (
	"net/http"
	"net/http/httptest"
)

type endpointConfig struct {
	Path           string
	ResponseString string
	ResponseHeader http.Header
}

type serverConfig struct {
	Endpoints []endpointConfig
}

func CreateTestServer(config serverConfig) *httptest.Server {
	mux := http.NewServeMux()

	for _, endpoint := range config.Endpoints {
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
