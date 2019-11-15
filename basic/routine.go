package main

import (
	"log"
	"time"
)

// Note that the log package is safe to use concurrently,
// but the fmt package is not.

func countDown() {
	for i := 5; i >= 0; i-- {
		log.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// Kick off a thread
	go countDown()

	// Since functions are first-class
	// you can write an anonymous function
	// for goroutine
	go func() {
		time.Sleep(time.Second * 2)
		log.Println("Delayed greetings!")
	}()

	// Use channels to signal when complete
	// or in this case just wait
	time.Sleep(time.Second * 4)
}
