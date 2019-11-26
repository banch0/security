package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

var ipToScan = "127.0.0.1"
var minPort = 0
var maxPort = 1024

func main() {
	activeTreads := 0
	doneChannel := make(chan bool)

	for port := minPort; port <= maxPort; port++ {
		go testTCPConnection(ipToScan, port, doneChannel)
		activeTreads++
	}
	// Wait for all threads to finish
	for activeTreads > 0 {
		<-doneChannel
		activeTreads--
	}
}

func testTCPConnection(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err == nil {
		log.Printf("Port %d: Open\n", port)
	}
	doneChannel <- true
}
