package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create hard link
	// editing content of one affected to another
	// Deleting or reneming one will not affect

	err := os.Link("test.txt", "test_link.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Symlink("test.txt", "test_symlink.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := os.Lstat("test_symlink.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Link info : %+v", fileInfo)

	// Change fileownership  of a symlink only
	// and not the file it points to
	err = os.Lchown("test_symlink.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}
