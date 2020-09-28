package main

import (
	"fmt"
	"net/http"
)

func firstMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("First middleware beforing calling next.")
		next.ServeHTTP(w, r)
		fmt.Println("First middleware after calling next.")
	})
}

func secondMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Second middleware beforing calling next.")
		next.ServeHTTP(w, r)
		fmt.Println("Second middleware after calling next.")
	})
}
