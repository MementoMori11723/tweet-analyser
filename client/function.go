package client

import (
	"html/template"
	"net/http"
)

var (
  path = "client/pages/"
)

func home(w http.ResponseWriter, r *http.Request) {
  temp, err := template.ParseFiles(
    path + "index.html",
  )

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := temp.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
