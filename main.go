// This file handles the application entry
// configuration and setup

package main

import (
	"goblog/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Index)
	server := http.Server{
		Addr:              "127.0.0.1:2000",
		Handler:           nil,
	}
	server.ListenAndServe()
}
