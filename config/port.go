package config

import (
  "os"
)

var (
  PORT string
)

func init() {
  PORT = os.Getenv("PORT")
  if PORT == "" {
    PORT = "8080"
  }
}

func GetPort() string {
  return PORT
}
