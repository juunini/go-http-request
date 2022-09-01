package http_request

import (
	"io"
	"net/http"
)

func Request(method, url string, headers *map[string]string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for key, value := range *headers {
			request.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	return response, err
}
