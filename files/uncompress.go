package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	gzipFile, err := os.Open("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}

	// Create  a gzip reader on top of the file reader
	// Again, it could be any type reader though
	gzipFile, err = gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipFile.Close()

	//Uncompress, to a writer. We'll use a file writer
	outfileWriter, err := os.Create("unzipped.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer outfileWriter.Close()

	// Copy contents of gzipped file to output file
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		log.Fatal(err)
	}

}
