package errors

import (
	"encoding/json"
	"strings"
)

type DBError struct {
	Status  int
	Code    string
	Message string
}

func (err *DBError) Error() string {
	b, _ := json.Marshal(err)
	return string(b)
}

func IsRecordNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "record not found")
}
