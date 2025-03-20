package config

import (
	"log"
	"os"
)

var (
	PORT        string
	GPT_API_KEY string
	GPT_URL     = "https://api.deepseek.com/chat/completions"
)

func init() {
	GPT_API_KEY = os.Getenv("API_KEY")
	if GPT_API_KEY == "" {
		log.Fatal("API_KEY not found")
	}
	PORT = os.Getenv("PORT")
	if PORT == "" {
    PORT = "8080"
	}
  PORT = ":" + PORT
}

func GetPort() string {
	return PORT
}

func GetGptData() (string, string) {
	return GPT_API_KEY, GPT_URL
}
