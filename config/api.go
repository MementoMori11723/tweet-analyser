package config

import (
	"log"
	"os"
)

var (
  GPT_API_KEY string
  GPT_URL = "https://api.openai.com/v1/chat/completions"
)

func init() {
  GPT_API_KEY = os.Getenv("API_KEY")
  
  if GPT_API_KEY == "" {
    log.Fatal("API_KEY not found")
  }
}

func GetGptData() (string, string) {
  return GPT_API_KEY, GPT_URL
}
