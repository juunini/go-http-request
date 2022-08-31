# Go http request

## Install

```sh
$ go get -u github.com/juunini/go-http-request
```

## Usage

```go
import http_request "github.com/juunini/go-http-request"

// ...

response, err := http_request.Request(
  method,
  url,
  headers,
  http_request.MapToRequestBody(body),
)

var body map[string]interface{}
err := http_request.ParseBody(response.Body, &body)

// ...
```

## API

```go
func Request(method, url string, headers *map[string]string, body io.Reader) (*http.Response, error)

func ParseBody[T any](body io.ReadCloser, target T) error

func MapToRequestBody(input any) *bytes.Reader

func StringToRequestBody(input string) *bytes.Reader
```
