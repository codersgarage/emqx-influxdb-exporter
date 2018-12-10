package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type object map[string]interface{}

type response struct {
	Code   string      `json:"code,omitempty"`
	Status int         `json:"-"`
	Title  string      `json:"title,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Errors error       `json:"errors,omitempty"`
}

func (r *response) ServerJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		panic(err)
	}
}

func parseBody(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func forward(w http.ResponseWriter, resp *http.Response) {
	defer resp.Body.Close()

	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
