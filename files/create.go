package main

import (
	"log"
	"os"
)

func main() {
	// Creating the file
	data, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)

	// Stat returns file info. It will return
	// an error if there is no file.
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File name: ", fileInfo.Name())
	log.Println("File size in bytes: ", fileInfo.Size())
	log.Println("File permission: ", fileInfo.Mode())
	log.Println("Last modified: ", fileInfo.ModTime())
	log.Println("Is Directory: ", fileInfo.IsDir())
	log.Printf("System interface type: %T\n", fileInfo.Sys())
	log.Printf("System info: %+v\n\n", fileInfo.Sys())

	defer data.Close()
}

// For renaming the file name use os.Rename()
// originalPath := "test.txt"
// newPath := "test2.txt"
// err := os.Rename(originalPath, newPath)

// For deleting the file use os.Remove()
