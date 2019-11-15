package main

import (
	"log"
	"time"
)

// More about chan philosophy
// https://blog.golang.org/share-memory-by-communicating

// Do some processing that takes a long time
// in a separate thread and signal when done
func process(doneChannel chan bool) {
	time.Sleep(time.Second * 3)
	doneChannel <- true
}

func main() {
	// Each channel can support one data type.
	var doneChannel chan bool
	// Channels are nil until initialized with make
	doneChannel = make(chan bool)

	// Kick off a lengthly process that will
	// signal when complete
	go process(doneChannel)

	// Get the first bool available in the channel
	// This is a blocking operation so executing
	// will not progress until value is recieved
	tempBool := <-doneChannel
	log.Println(tempBool)

	// or to simply ignore the value but still wait
	// <-doneChannel

	// Start another process thread to run it background
	// and signal when done
	go process(doneChannel)

	// Make channel non-blocking with select statement
	// This gives you the abality to continue executing
	// even if no message waiting in the channel
	var readyToExit = false
	for !readyToExit {
		select {
		case done := <-doneChannel:
			log.Println("Done message recieved.", done)
			readyToExit = true
		default:
			log.Println("No done signal yet. Waiting.")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
