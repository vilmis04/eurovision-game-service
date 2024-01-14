package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeJson(req *http.Request, body interface{}) error {
	return json.NewDecoder(req.Body).Decode(body)
}
