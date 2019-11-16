package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create output file
	newFile, err := os.Create("devdungeon.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// HTTP Get request
	url := "http://www.devdungeon.com/archive"
	response, err := http.Get(url)
	defer response.Body.Close()

	newBytesWriten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", newBytesWriten)
}
