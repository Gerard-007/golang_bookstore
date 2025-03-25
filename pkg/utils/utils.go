package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(resp *http.Response, x interface{}) error {
	if body, err := io.ReadAll(resp.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return fmt.Errorf("failed to unmarshal JSON: %w", err)
		}
	}
	return nil
}
