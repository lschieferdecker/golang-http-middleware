package main

import (
	"fmt"
	"github.com/lschieferdecker/golang-http-middleware/real/middleware"
	"net/http"
)

func main() {
	hello := http.HandlerFunc(helloHandler)
	http.Handle("/hello", middleware.Authenticate(middleware.Validate(hello)))

	bye := http.HandlerFunc(byeHandler)
	http.Handle("/bye", middleware.Authenticate(middleware.Validate(bye)))

	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello :)")
}

func byeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Bye :)")
}
