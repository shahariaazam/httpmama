package httpmama

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCreateTestServer(t *testing.T) {
	endpoint1 := endpointConfig{
		Path:           "/foo",
		ResponseString: "hello, world!",
		ResponseHeader: http.Header{"Content-Type": []string{"text/plain"}},
	}

	endpoint2 := endpointConfig{
		Path:           "/bar",
		ResponseString: "goodbye, world!",
		ResponseHeader: http.Header{"Content-Type": []string{"text/plain"}},
	}

	config := serverConfig{
		Endpoints: []endpointConfig{endpoint1, endpoint2},
	}

	server := CreateTestServer(config)
	defer server.Close()

	resp1, err := http.Get(server.URL + "/foo")
	if err != nil {
		t.Errorf("error making GET request: %v", err)
	}
	defer resp1.Body.Close()

	body1, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}

	if string(body1) != "hello, world!" {
		t.Errorf("expected response body to be 'hello, world!', got '%s'", string(body1))
	}

	resp2, err := http.Get(server.URL + "/bar")
	if err != nil {
		t.Errorf("error making GET request: %v", err)
	}
	defer resp2.Body.Close()

	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}

	if string(body2) != "goodbye, world!" {
		t.Errorf("expected response body to be 'goodbye, world!', got '%s'", string(body2))
	}
}
