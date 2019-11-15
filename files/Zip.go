package main

import (
	"archive/zip"
	"log"
	"os"
)

func main() {
	// Create a file to write the archive buffer to
	// Could also use an in memory buffer
	outFile, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Creates a zip writer on top of the file writer
	zipWriter := zip.NewWriter(outFile)

	// Create struct and adding data
	var fileToArchive = []struct {
		Name, Body string
	}{
		{"test.txt", "Strings of contenst"},
		{"test2.txt", "\x61\x62\x63\n"},
	}

	// append data to archive
	for _, file := range fileToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Clean up
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
