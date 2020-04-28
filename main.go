// This file handles the application entry
// configuration and setup

package main

import (
	"github.com/julienschmidt/httprouter"
	"goblog/route"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	route.Route(router)
	server := &http.Server{
		Addr:    "127.0.0.1:2000",
		Handler: router,
	}
	log.Printf("Server started at http://%s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
