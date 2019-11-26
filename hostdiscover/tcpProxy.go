package main

import (
	"log"
	"net"
)

var localListenAddress = "localhost:9999"
var remoteHostAddress = "localhost:3000" // Not required to be remote
func main() {
	listener, err := net.Listen("tcp", localListenAddress)
	if err != nil {
		log.Fatal("Error creating listener. ", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection. ", err)
		}
		go handleConnection(conn)
	}
}

// Forward the request to the remote host and pass response
// back to client
func handleConnection(localConn net.Conn) {
	// Create remote connection that will recieve forwarded data
	remoteConn, err := net.Dial("tcp", remoteHostAddress)
	if err != nil {
		log.Fatal("Error creating listener. ", err)
	}
	defer remoteConn.Close()

	// Read from the client and forward to remote host
	buf := make([]byte, 4096) // 4k buffer
	numBytesRead, err := localConn.Read(buf)
	if err != nil {
		log.Println("Error reading from client. ", err)
	}
	log.Printf("Forwarding from %s to %s: \n%s\n\n",
		localConn.LocalAddr(),
		remoteConn.RemoteAddr(),
		buf[0:numBytesRead])
	_, err = localConn.Write(buf[0:numBytesRead])
	if err != nil {
		log.Println("Error writing back to client. ", err)
	}
}
