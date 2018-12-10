package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func ParseBody(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func ParseInt(s string) (int, error) {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}
