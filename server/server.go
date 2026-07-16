package server

import (
	"log"
	"net"
)

func StartSever() {
	address := "localhost:8080"

	listener, err := net.Listen("tcp",address)
	if err != nil {
		log.Fatalf("Couldn't set up the connection!")
	}

	defer listener.Close()

	log.Printf("Server is ready to receive requests on %s", address)

	// in a conversation
	for {
		conn, err := listener.Accept()
		if err != nil {
			// didn't accept just go to another one 
			continue
		}

		// to a client 
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	log.Printf("New Client connected: %s ", conn.RemoteAddr())

	// start conversation
	buffer := make([]byte, 1024)

	for {
		received, err := conn.Read(buffer)
		if err != nil {
			// didn't receive => return
			log.Printf("Connection error: %v", err)
			return
		}

		message := string(buffer[:received])		// only received bytes
		log.Printf("Received: %s", message)
		conn.Write([]byte(message))
	}
}