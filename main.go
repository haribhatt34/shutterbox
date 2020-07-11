package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// Set Content type to indicate the browser how to render page,
	// if text/plain is used, it won't parse html tags, and out page as it is.
	// w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Type", "text/html")

	// print to fresh console
	fmt.Println("Someone visited our website")
	fmt.Fprint(w, "<h1>Welcome to my Awesome site!</h1>")
}

func main() {
	// maps the root to a method
	http.HandleFunc("/", handlerFunc)
	// starts the server
	http.ListenAndServe(":3000", nil)
}
