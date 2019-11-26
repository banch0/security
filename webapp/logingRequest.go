package main

import (
	"fmt"
	"github.com/urfave/negroni"
	"net/http"
)

// Request response to client
func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "You requested: "+req.URL.Path)
}

func main() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", indexHandler)

	negroniHandler := negroni.New()
	negroniHandler.Use(negroni.NewLogger()) // Negroni's default logger
	negroniHandler.UseHandler(multiplexer)

	http.ListenAndServe("localhost:3000", negroniHandler)
}
