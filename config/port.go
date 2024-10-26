package config

import (
  "github.com/joho/godotenv"
  "os"
)

var (
  PORT string
)

func init() {
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }
  PORT = os.Getenv("PORT")
  if PORT == "" {
    PORT = "8080"
  }
}

func GetPort() string {
  return PORT
}
