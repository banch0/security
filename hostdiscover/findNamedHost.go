package main

/*
The pure Go resolver is default and can only block a goroutine instead of a
system thread, making it a little more efficient. You can explicitly set the DNS
resolver with an environment variable:
export GODEBUG=netdns=go
export GODEBUG=netdns=cgo
# Use pure Go resolver (default)
# Use cgo resolver
*/
import (
	"log"
	"net"
	"strconv"
	"strings"
)

var subnetToScan = "192.168.0" // First three octets

func main() {
	activeThreads := 0
	doneChannel := make(chan bool)

	for ip := 0; ip <= 255; ip++ {
		fullIP := subnetToScan + "." + strconv.Itoa(ip)
		go resolve(fullIP, doneChannel)
		activeThreads++
	}

	// Wait for all threads to finish
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func resolve(ip string, doneChannel chan bool) {
	addresses, err := net.LookupAddr(ip)
	if err != nil {
		log.Printf("%s - %s\n", ip, strings.Join(addresses, ", "))
	}
}
