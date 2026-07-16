package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func StartClient() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Printf("Couldn't connect to the server: %v", err)
		return
	}

	defer conn.Close()

	fmt.Println("Connected to server. Start chatting!")

	// send from own 
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		if !scanner.Scan() {
			// no input  | EOF
			break		
		}

		message := scanner.Text()		// in string
		if message == "exit" {
			break		// end the conversation
		}

		// active conversation => wait for the server's response
		respone := make([]byte, 1024)
		received, err := conn.Read(respone)
		if err != nil {
			// didn't received 
			log.Printf("Error receiving the response: %v", err)
			break		// end 
		}

		fmt.Println(string(respone[:received]))
	}
}