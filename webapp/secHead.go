package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

// Custom middleware handler logs user agent
func addSecureHeaders(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	res.Header().Add("Content-Security-Policy", "default-src, 'self'")
	res.Header().Add("X-Frame-Options", "SAMEORIGIN")
	res.Header().Add("X-XSS-Protection", "1; mode=block")
	res.Header().Add("Strict-Transport-Security", "max-age=10000, includeSubdomains; preloaded")
	res.Header().Add("X-Content-Type-Options", "nosniff")
	next(res, req) // Pass on to next middleware handler
}

// Return response to client
func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "You requested: "+req.URL.Path)
}

func main() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", indexHandler)

	negroniHandler := negroni.New()

	// Set up as many middleware functions as you need, in order
	negroniHandler.Use(negroni.HandlerFunc(addSecureHeaders))
	negroniHandler.Use(negroni.NewLogger())
	negroniHandler.UseHandler(multiplexer)

	http.ListenAndServe("locahost:3000", negroniHandler)
}
