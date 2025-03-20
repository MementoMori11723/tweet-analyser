package server

import (
	"net/http"
)

var routes = map[string]func(http.ResponseWriter, *http.Request){
	"POST /api": api,
	"GET /":     home,
}

func New() *http.ServeMux {
	mux := http.NewServeMux()
	for path, handler := range routes {
		mux.HandleFunc(path, handler)
	}
	return mux
}
