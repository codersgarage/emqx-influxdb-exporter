package api

import (
	"net/http"
)

func recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {

				resp := response{}
				resp.Title = "Something went wrong"
				resp.Errors = rvr.(error)
				resp.Status = http.StatusInternalServerError
				resp.ServerJSON(w)
				return
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
