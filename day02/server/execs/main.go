package main

import (
	"demos"
	"net/http"
)

func main() {
	http.HandleFunc("/req/get", demos.DealGetHandler)
	http.HandleFunc("/req/post", demos.DealPostHandler)
	http.ListenAndServe(":8081", nil)
}
