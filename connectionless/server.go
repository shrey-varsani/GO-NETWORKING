package connectionless

import (
	"fmt"
	"log"
	"net"
)

// ListenUDP and WriteTo

func main() {

	// set up the UDP address
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Fatal("Couldn't resolve address:", err)
	}

	// start listening 
	conn, err := net.ListenUDP("udp", addr)			// same as tcp => Listen 
	if err != nil {
		log.Fatal("Listen failed:", err)
	}

	defer conn.Close()

	// buffer for incoming data
	buffer := make([]byte, 1024)
	for {
		// get clientAddr without conn.RemoteAddr() (unlike tcp)
		received, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Read error: %v", err)
			continue
		}

		fmt.Printf("Got message from %s: %s\n", clientAddr, string(buffer[:received]))		// bytes => string

		// send back to client
		_, err = conn.WriteToUDP(buffer[:received], clientAddr)
		if err != nil {
			log.Printf("Write error: %v", err)
		}
	}
}