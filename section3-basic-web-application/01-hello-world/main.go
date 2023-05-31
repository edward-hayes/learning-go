package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<H1>Hello, World!</>")
	})

	http.ListenAndServe(":8080", nil)
}
