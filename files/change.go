package main

import (
	"log"
	"os"
	"time"
)

func main() {
	// Change the permission using Linuxs style
	err := os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	// Change ownership
	err = os.Chown("test.txt", os.Geteuid(), os.Getegid())
	if err != nil {
		log.Println(err)
	}

	// Change timestamps
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}
