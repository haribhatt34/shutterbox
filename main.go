package main

import (
	"fmt"
	"net/http"
)

/*	Notes
 *
 * printf(...) is equivalent to fprintf(stdout,...).
 * fprintf is used to output to stream.
 * sprintf(buffer,...) is used to format a string to a buffer.
 *
 */

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// Set Content type to indicate the browser how to render page,
	// if text/plain is used, it won't parse html tags, and out page as it is.
	// w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Type", "text/html")

	// print to fresh console
	fmt.Println("Someone visited our website")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my Awesome site!</h1>")
	} else if r.URL.Path == "/contact" || r.URL.Path == "/contact/" {
		// we can also use `` around string, to spread it into multiple lines.
		fmt.Fprint(w, "To get in touch, please send us a mail to <a href=\"mailto:shutterbox@gmail.com\"> shutterbox@gmailcom </a>.")
		fmt.Fprint(w, "<br><br>")
		fmt.Fprint(w, "Thanks For Visiting!!!")
	} else {
		// If WriteHeader is not called explicitly, the first call to Write
		// will trigger an implicit WriteHeader(http.StatusOK).
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1> 404, Page Not found !!!")
	}

	// format the output
	//fmt.Fprintf(w, "<p>This is your current path: %v", r.URL.Path)
}

func main() {
	// A small Caveat, HandleFunc always matches the largest path,
	// irrespective of the order, i.e. which mapping if defined early.
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", mux)
	// maps the root to a method
	//http.HandleFunc("/", handlerFunc)
	// starts the server
	//http.ListenAndServe(":3000", nil)
}
