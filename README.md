<h1 align="center">HTTP Mama</h1>
<p align="center">Utility library for various HTTP related task in Golang</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/shahariaazam/httpmama"><img src="https://pkg.go.dev/badge/github.com/shahariaazam/httpmama.svg" height="20"/></a>
</p><br/><br/>

<p align="center">
  <a href="https://github.com/shahariaazam/httpmama/actions/workflows/CI.yaml"><img src="https://github.com/shahariaazam/httpmama/actions/workflows/CI.yaml/badge.svg" height="20"/></a>
  <a href="https://codecov.io/gh/shahariaazam/httpmama"><img src="https://codecov.io/gh/shahariaazam/httpmama/branch/master/graph/badge.svg?token=NKTKQ45HDN" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_httpmama"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_httpmama&metric=reliability_rating" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_httpmama"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_httpmama&metric=vulnerabilities" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_httpmama"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_httpmama&metric=security_rating" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_httpmama"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_httpmama&metric=sqale_rating" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_httpmama"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_httpmama&metric=code_smells" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_httpmama"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_httpmama&metric=ncloc" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_httpmama"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_httpmama&metric=alert_status" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_httpmama"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_httpmama&metric=duplicated_lines_density" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_httpmama"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_httpmama&metric=bugs" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_httpmama"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_httpmama&metric=sqale_index" height="20"/></a>
</p><br/><br/>

## Usage

```shell
go get github.com/shahariaazam/httpmama
```

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shahariaazam/httpmama"
)

func main() {
	endpoint1 := httpmama.TestEndpoint{
		Path:           "/foo",
		ResponseString: "hello, world!",
		ResponseHeader: http.Header{"Content-Type": []string{"text/plain"}},
	}

	endpoint2 := httpmama.TestEndpoint{
		Path:           "/bar",
		ResponseString: "goodbye, world!",
		ResponseHeader: http.Header{"Content-Type": []string{"text/plain"}},
	}

	config := httpmama.ServerConfig{
		TestEndpoints: []httpmama.TestEndpoint{endpoint1, endpoint2},
	}

	// Create the server
	server := httpmama.NewTestServer(config)
	defer server.Close()

	// Start making request to the server
	resp, err := http.Get(server.URL + "/foo")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body)) // Output: "hello, world!"
}

```

## Documentation

Full documentation is available on [pkg.go.dev/github.com/shahariaazam/httpmama](https://pkg.go.dev/github.com/shahariaazam/httpmama#section-documentation)


### 📝 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/shahariaazam/httpmama/blob/master/LICENSE) file for details.