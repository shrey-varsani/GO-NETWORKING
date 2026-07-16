package main

import (
	"Go-Networking/client"
	"Go-Networking/server"
	"time"
	// "net/http"
)

func main()  {
	go func() {
		server.StartSever()
	}()

	time.Sleep(time.Second)		// start everything

	client.StartClient()
} 