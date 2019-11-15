package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// buffer lets you create a buffered writer so that
	// you can work with a buffer in memory before
	// writing it to dist

	// Open file for writing
	file, err := os.OpenFile("test_link.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a buffered writer from the file
	bufferedWriter := bufio.NewWriter(file)

	// Write bytes to buffer
	bytesWritten, err := bufferedWriter.Write(
		[]byte{65, 66, 67},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes writen: %d\n", bytesWritten)

	// Checking how much is stored in buffer waiting
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// See how much buffer is available
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// write memory buffer to disk
	bufferedWriter.Flush()

	bufferedWriter.Reset(bufferedWriter)

	// See how much buffer is available
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// Change buffer size
	bufferedWriter = bufio.NewWriterSize(
		bufferedWriter,
		8000,
	)

	// Check available buffer size after resizing
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d \n", bytesAvailable)
}
