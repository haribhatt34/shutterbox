package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my Awesome site!</h1>")
}

func main() {
	// maps the root to a method
	http.HandleFunc("/", handlerFunc)
	// starts the server
	http.ListenAndServe(":3000", nil)
}
