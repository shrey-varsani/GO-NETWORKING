package main

import (
	// "Go-Networking/client"
	// "Go-Networking/server"
	"Go-Networking/unixSockets"
	"fmt"
	// "log"
	"net"
	"sync"
	"time"
	// "time"
	// "net/http"
)

func handleClient(conn *net.UDPConn, wg *sync.WaitGroup, clientAddr *net.UDPAddr, data []byte) {
	defer wg.Done()

	fmt.Printf("Handling data: %s from client %s", string(data), clientAddr)

	_, err := conn.Write(data)
	if err != nil {
		fmt.Printf("error while sending the data: %v", err)
	}
}

func main()  {
	go unixsockets.StartServer()

	time.Sleep(time.Second)

	unixsockets.StartClient()
	
	// go func() {
	// 	server.StartSever()
	// }()
	// time.Sleep(time.Second)		// start everything
	// client.StartClient()

	// addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	// if err != nil {
	// 	log.Fatalf("Couldn't resolve the address: %v", err)
	// }

	// conn, err := net.ListenUDP("udp", addr)
	// if err != nil {
	// 	log.Fatalf("Couldn't listen to connection: %v", err)
	// }

	// defer conn.Close()

	// buffer := make([]byte, 1024)
	// var wg sync.WaitGroup

	// for {
	// 	// read from the client first 
	// 	n, clientAddr, err := conn.ReadFromUDP(buffer)
	// 	if err != nil {
	// 		log.Printf("Read Error: %v", err)
	// 		continue 
	// 	}

	// 	wg.Add(1)
	// 	go handleClient(conn, &wg, clientAddr, buffer[:n])
	// }
	
} 