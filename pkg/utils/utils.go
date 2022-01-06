package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody parses the post request body into the specified struct variable type
func ParseBody(r *http.Request, x interface{}) {

	if body, err := io.ReadAll(r.Body); err == nil {

		if err := json.Unmarshal([]byte(body), x); err == nil {
			return
		}
	}
}
