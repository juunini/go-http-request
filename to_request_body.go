package http_request

import (
	"bytes"
	"encoding/json"
)

func MapToRequestBody(input any) *bytes.Reader {
	body, _ := json.Marshal(input)

	return bytes.NewReader(body)
}

func StringToRequestBody(input string) *bytes.Reader {
	return bytes.NewReader([]byte(input))
}
