package main

import (
	"fmt"
	"net/http"
)

// hello and headers are handlers:
// handler is an object implementing the http.Handler interface.
// To write a handler is by using the http.HandlerFunc adapter on functions with the appropriate signature.
// Functions serve as handlers take a http.ResponseWriter and a http.Request as arguments.
// The response writer is used to fill in the HTTP response. which is "hello\n" in this case.
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// Register handlers on server routes using http.HandleFunc function.
	// This setes the default router in the net/http package and takes a function as an argument.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	//nil tells it to use the default router set up above.
	http.ListenAndServe(":8090", nil) // When the second param is nil, DefaultServerMux will be used.
}
