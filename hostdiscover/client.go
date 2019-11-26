package main

import (
	"log"
	"net"
)

var protocol = "tcp" // tcp or udp
var remoteHostAddress = "localhost:3000"

func main() {
	conn, err := net.Dial(protocol, remoteHostAddress)
	if err != nil {
		log.Fatal("Error creating listener. ", err)
	}
	conn.Write([]byte("Hello, server. Are you there?"))

	serverResponseBuffer := make([]byte, 4096)
	numBytesRead, err := conn.Read(serverResponseBuffer)
	if err != nil {
		log.Print("Error reading from server: ", err)
	}
	log.Println("Message recieved from server:")
	log.Printf("%s\n", serverResponseBuffer[0:numBytesRead])
}
