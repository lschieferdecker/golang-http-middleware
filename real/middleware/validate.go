package middleware

import (
	"fmt"
	"net/http"
)

func Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		if len(queryParams) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Validation failed.")
			return
		}
		next.ServeHTTP(w, r)
	})
}
