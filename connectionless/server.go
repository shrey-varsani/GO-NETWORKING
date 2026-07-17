package connectionless

import (
	"fmt"
	"log"
	"net"
)

func StartServer() {
	// set udp adderss

	addr, err := net.ResolveUDPAddr("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Couldn't resolve the address", err)
	}

	// listen 
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("Failed to listen", err)
	}

	defer conn.Close()

	// incoming data
	buffer := make([]byte, 1024)
	for {
		n,_, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Read error: %v", err)
			continue
		}

		// read => string 
		fmt.Printf("Received message: %s from %s", string(buffer[:n]), addr)

		// send back 
		_, err = conn.WriteToUDP(buffer[:n], addr)
		if err != nil {
			log.Printf("Write error: %v", err)
		}
	}
}