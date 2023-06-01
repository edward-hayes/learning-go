package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>This is the home page</>")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>About Me</>")
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(portNumber, nil)
}
