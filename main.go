// This file handles the application entry
// configuration and setup

package main

import (
	"goblog/handler"
	"goblog/handler/user"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", handler.Index)
	mux.HandleFunc("/signup", user.New)
	mux.HandleFunc("/signup_account", user.Create)
	mux.HandleFunc("/users", user.Index)
	server := &http.Server{
		Addr:    "127.0.0.1:2000",
		Handler: mux,
	}
	log.Printf("Server strted at %s", server.Addr)
	err := server.ListenAndServe()
	if err == http.ErrServerClosed || err == nil {
		log.Println("Server closed", err)
	}
}
