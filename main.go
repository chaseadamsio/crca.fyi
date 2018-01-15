package main

import (
	"fmt"
	"net/http"
)

var redirects = map[string]string{
	"foo": "https://chaseadams.io",
}

func handler(wtr http.ResponseWriter, req *http.Request) {
	from := req.URL.Path[1:]
	if to, ok := redirects[from]; ok {
		http.Redirect(wtr, req, to, http.StatusPermanentRedirect)
	}
	fmt.Fprintf(wtr, "Hi there I love %s", req.URL.Path[1:])
}

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
