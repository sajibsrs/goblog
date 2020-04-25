// This file handles the application entry
// configuration and setup

package main

import (
	"github.com/julienschmidt/httprouter"
	"goblog/handler"
	"goblog/handler/user"
	"log"
	"net/http"
)

func main() {
	mux := httprouter.New()
	mux.ServeFiles("/static/*filepath", http.Dir("/home/sajib/playground/goblog/static"))
	mux.GET("/", handler.Index)
	mux.POST("/signup", user.New)
	mux.GET("/signup", user.New)
	mux.GET("/users", user.Index)
	server := &http.Server{
		Addr:    "127.0.0.1:2000",
		Handler: mux,
	}
	log.Printf("Server started at %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
