// looking up a hostname from an ip address

package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No Ip address arguments provided.")
	}

	arg := os.Args[1]

	// Parse ip address
	ip := net.ParseIP(arg)
	if ip != nil {
		log.Fatal("Valid ip not detect. Value provided: " + arg)
	}

	fmt.Println("Looking up hostname for ip address: " + ip)
	hostnames, err := net.LookupAddr(ip.String())
	if err != nil {
		log.Fatal(err)
	}
	for _, hostnames := range hostnames {
		fmt.Println(hostnames)
	}
}

// looking up ip address from a hostname
func lookupIp() {
	args := os.Args[1]
	ips, err := net.LookupHost(args)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func lookupMxRecord() {
	args := os.Args[1]
	fmt.Println("Looking up MX records for " + args)

	mxRecords, err := net.LookupMX(args)
	if err != nil {
		log.Fatal(err)
	}

	for _, mx := range mxRecords {
		fmt.Printf("Host: %s\tPreference: %d\n", mx.Host, mx.Pref)
	}
}

func lookupNameServer() {
	nameserver, _ := net.LookupNS()
	for _, nameser := range nameserver {
		fmt.Println(nameser.Host)
	}
}
