package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pygola/webdemos/hello"
)

func main() {
	muxx := http.NewServeMux()
	muxx.HandleFunc("/", hello.HelloWorld)

	fmt.Println("http://localhost:8080")
	s := &http.Server{
		Addr:           ":8080",
		Handler:        muxx,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

/*

ServeMux
The ServeMux is a multiplexor (or simply an HTTP request router) that compares incoming HTTP requests
against a list of predefined URI resources and then calls the associated handler for the resource requested by
the HTTP client.
Handler

The ServeMux provides a multiplexor and calls corresponding handlers for HTTP requests. Handlers are
responsible for writing response headers and bodies. In Go, any object can become a handler, thanks to Go’s
excellent interface implementation provided by its type system. If any object satisfies the implementation of
the http.Handler interface, it can be a handler for serving HTTP requests.


type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/
