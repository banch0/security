package main

import (
	"io"
	"log"
	"os"
)

// The nameOfFile.Seek() func is useful for setting
// cursor in specified location

func main() {
	// Open original file
	originalFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// Create new file
	copyOfFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer copyOfFile.Close()

	// Copy the bytes destination from source
	bytesWritten, err := io.Copy(copyOfFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	// Commit the file contents
	// Flushes memory to disk
	err = copyOfFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
