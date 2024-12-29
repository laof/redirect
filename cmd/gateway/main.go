package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/feed", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not found: ")
}
