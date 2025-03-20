package main

import (
	"log"
	"net/http"
	"os"
	"tweet-analyser/config"
	"tweet-analyser/server"
)

func main() {
	PORT := config.GetPort()
	server := http.Server{
		Addr:    PORT,
		Handler: server.New(),
	}
	log.Println("Server is running on http://localhost" + server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
