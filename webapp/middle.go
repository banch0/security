package main

import (
	"fmt"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

// Custom middleware handler logs user agent
func customMiddlewareHandler(res http.ResponseWriter,
	req *http.Request,
	next http.HandlerFunc) {
	log.Println("Incoming request: " + req.URL.Path)
	log.Println("User agent: " + req.UserAgent())

	next(res, req) // Pass on to next middleware handler
}

// Return response to client
func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "You requested:"+request.URL.Path)
}

func main() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", indexHandler)

	negroniHandler := negroni.New()
	negroniHandler.Use(negroni.HandlerFunc(customMiddlewareHandler))
	negroniHandler.UseHandler(multiplexer)
	http.ListenAndServe("localhost:3000", negroniHandler)
}
