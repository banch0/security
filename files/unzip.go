package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Create a reader out of the zip archive
	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		log.Fatal(err)
	}

	defer zipReader.Close()

	// Iterate through each file/dir found in
	for _, file := range zipReader.Reader.File {
		// open the file inside the zip archive
		// like a normal file
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		if file.FileInfo().IsDir() {
			// Create directories to recreate directory
			// structure inside a zip archive. Also
			// preserves permissions
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			// Extract regular file since not a directory
			log.Println("Extracting file:", file.Name)

			// open an output file for writing
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			// "Extract the file by copying zipped file"
			// contents to the output file
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}

		}
	}
}
