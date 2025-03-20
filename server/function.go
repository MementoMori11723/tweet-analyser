package server

import (
	"embed"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"tweet-analyser/server/gpt"
	"tweet-analyser/server/scraper"
)

var (
  //go:embed pages/*
  files embed.FS
  indexTmpl *template.Template
)

func init() {
  var err error
  indexTmpl, err = template.ParseFS(files, "pages/index.html")
  if err != nil {
    log.Println(err.Error())
    os.Exit(1)
  }
}

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

func home(w http.ResponseWriter, _ *http.Request) {
  if err := indexTmpl.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } 
}
