package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"tweet-analyser/server/gpt"
	"tweet-analyser/server/scraper"
)

func api(w http.ResponseWriter, r *http.Request) {
  log.Println("API called")
	url, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("URL: ", string(url))
	gptData := scraper.New(string(url))
  log.Println("sending to GPT")
  message := gpt.New(gptData)
  log.Println("recived response from GPT")
	data := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
