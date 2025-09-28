package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет, мир!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("HTTP server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
