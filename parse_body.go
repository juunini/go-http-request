package http_request

import (
	"encoding/json"
	"io"
)

func ParseBody[T any](body io.ReadCloser, target T) error {
	reader, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(reader, target)
}
