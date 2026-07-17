package connectionless

import (
	"fmt"
	"log"
	"net"
	"time"
)

func StartClient() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil{
		log.Fatal("Couldn't resolve address:", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Couldn't Dial the address", err)
	}

	defer conn.Close()

	// start conversation
	message := []byte("Hello, UDP!")
	_, err = conn.Write(message)
	if err != nil {
		log.Printf("Send failed: %v", err)
		return
	}

	// wait for server reply with timelimit
	conn.SetDeadline(time.Now().Add(time.Second * 5))
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error while receiving: %v", err)
	}

	fmt.Printf("Server says: %s", string(buffer[:n]))
}