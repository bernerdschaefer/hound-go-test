package request

import "encoding/json"

// Request represents the data of an incoming request.
type Request struct {
}

// RequestFromJSON decodes a JSON payload into a Request.
func RequestFromJSON(data []byte) (Request, error) {
	var req Request

	if err := json.Unmarshal(data, &req); err != nil {
		return Request{}, err
	}

	return req, nil
}
