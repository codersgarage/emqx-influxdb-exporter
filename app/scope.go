package app

import (
	"net/http"
)

type Scope struct {
	Request *http.Request
}

func NewScope(r *http.Request) *Scope {
	return &Scope{
		Request: r,
	}
}
