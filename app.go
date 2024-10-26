package main

import (
	"fmt"
	"net/http"
	"tweet-analyser/client"
	"tweet-analyser/config"
	"tweet-analyser/server"
)

func main() {
	go func() {
		client := client.New()
		server := server.New()
		PORT := config.GetPort()

		http.Handle("/", client)
		http.Handle("/api", server)
    fmt.Println("Server is running on http://localhost:"+PORT)
    http.ListenAndServe(":"+PORT, nil)
	}()
	fmt.Scanln()
}
