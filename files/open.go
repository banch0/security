package main

import (
	"log"
	"os"
)

/*
	When calling os.Open(), it just requires filename and
	provides a read-only file

	Another option is to use os.OpenFile(), which expects more options.

*/

func main() {
	// Simple read-only open file.
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// Open file with more options.
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// Checking if file not exist
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Fatal does not exist.")
		}
	}
	log.Println("File does exist yes")
	log.Println(fileInfo)

	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("ERROR: Read permission denied")
		}
	}
	defer file.Close()

}

/*
	Use this attribures individually or combined
	with an OR for second arg of OpenFile()

	e.g os.O_CREATE|os.O_APPEND
	or os.O_CREATE|os.O_TRUNC|os.O_WRONLY

	os.O_RDONLY // read only
	os.O_WRONLY // write only
	os.O_RDWR // Read and write
	os.O_APPEND // append to end of file
	os.O_CREATE // Create if none exist
	os.O_TRUNC //Trucnate file when opening
*/
