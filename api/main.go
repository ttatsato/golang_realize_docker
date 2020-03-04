package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "sample")
	})
	log.Printf("Server running at http://localhost:%s/", "8080")
	http.ListenAndServe(":8080", nil)
}