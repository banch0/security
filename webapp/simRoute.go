package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	// Write contents of  the response body to the writer interface
	// Request object contains informantion about and from the client
	fmt.Printf(writer, "You requested: " + request.URL.Path)
}

func main() {
	http.HandleFunc("/", indexHandler)
	// err := http.ListenAndServe("localhost:8080", nil)
	err := http.ListenAndServeTLS(
		"localhost:8181",
		"cert.pem",
		"privateKey.pem",
		nil,
	)
	if err != nil {
		log.Fatal("Error creating servers", err)
	}
}
