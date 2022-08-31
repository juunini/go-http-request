# Go http request

## Install

```sh
$ go get -u github.com/juunini/go-http-request
```

## Usage

```go
import http_request "github.com/juunini/go-http-request"

// ...

response, err := http_request.Request(method, url, headers, body)

var body map[string]interface{}
if err := http_request.ParseBody(response.Body, &body); err != nil {
  panic(err)
}

// ...
```
