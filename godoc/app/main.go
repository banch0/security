package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Container runnings")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world!")
	})
	http.ListenAndServe(":8888", nil)
}
