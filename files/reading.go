package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*
	// Example read exactly n bytes

	bytesSlice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlice)

*/

func main() {
	// Open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		log.Println(err)
	}

	// os.File.Read(), io.ReadFull(), and
	// io.ReadAtLeast() all work with a fixed
	// byte slice that you make before you read

	// ioutil.ReadAll() will read every byte
	// from the render (in this case a file),
	// and return a slice of unknown slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))
}

/*
	// Similar to io.ReadAll() function io.ReadFile() will
	// read all the bytes in a file and return byte slice.

	// Read file to byte slice
	data, err := io.ReadFile("test.txt")
	log.Printf("Data read: %s\n", data)
*/

/*
	// Buffered reader
	Creating a buffered reader will store a memory buffer
	with some of the contents. A buffered reader also provides some
	more functions that are not available on the os.File or io.Reader types

	The default buffer size is 4096 and the minimum size is 16. Buffered
	readers provide a set of useful functions. Some of the functions
	available include, but are not limited to, the following:

	Read(): This is to read data into a byte slice
	Peek(): This is to inspect the next bytes without moving the file cursor
	ReadByte(): This is to read a single byte
	UnreadByte(): This unreads the last byte read
	ReadBytes(): This reads bytes until the specified delimeter is reached
	ReadString(): This reads a string until the specified delimeter is reached

	The following example demonstrate how to use buffered reader to get data
	from a file. First, it opens a file and then creates a buffered reader that
	wraps the file object.

	file, err := os.Open("test.txt")
	bufferedReader := bufio.NewReader(file)

	// Get bytes without advancing pointer
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(5)

	// Read and advance pointer
	numBytesRead, err := bufferedReader.Read(byteSlice)

	// Read 1 byte. Error if no byte to read
	myByte, err := bufferedReader.ReadByte()


	// Read up to and including delimiter
	// Returns byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')

	// Read up to and including delimiter
	// Returns string
	dataBytes, err := bufferedReader.ReadString('\n')

*/
