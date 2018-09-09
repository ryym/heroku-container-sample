package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Starting server")
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
}
