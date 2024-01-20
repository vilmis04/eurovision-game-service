package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeRequestJson(req *http.Request, body interface{}) error {
	return json.NewDecoder(req.Body).Decode(body)
}
