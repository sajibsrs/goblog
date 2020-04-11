// This file handles the application entry
// configuration and setup

package main

import (
	"goblog/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", handler.Index)
	mux.HandleFunc("/users/", handler.User)
	server := &http.Server{
		Addr:              "127.0.0.1:2000",
		Handler:           mux,
	}
	_ = server.ListenAndServe()
}
