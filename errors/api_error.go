package errors

import (
	"encoding/json"
)

type APIError struct {
	Status int    `json:"-"`
	Code   string `json:"-"`
	Title  string `json:"title"`
	Errors error  `json:"errors"`
}

func (err *APIError) Error() string {
	b, _ := json.Marshal(err)
	return string(b)
}

func NewAPIError(status int, code string, title string, errs error) *APIError {
	return &APIError{
		Status: status,
		Code:   code,
		Title:  title,
		Errors: errs,
	}
}
