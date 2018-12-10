package api

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	resp := response{
		Status: http.StatusOK,
		Data:   "OK",
	}
	resp.ServerJSON(w)
}
