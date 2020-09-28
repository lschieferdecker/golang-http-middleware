package middleware

import (
	"fmt"
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.Header["Token"]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Authentication failed.")
			return
		}
		next.ServeHTTP(w, r)
	})
}
