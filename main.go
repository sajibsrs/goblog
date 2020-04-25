// This file handles the application entry
// configuration and setup

package main

import (
	"github.com/julienschmidt/httprouter"
	"goblog/route"
	"goblog/route/user"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("/home/sajib/playground/goblog/static"))
	router.GET("/", route.Index)
	router.POST("/signup", user.New)
	router.GET("/signup", user.New)
	router.GET("/users", user.Index)
	server := &http.Server{
		Addr:    "127.0.0.1:2000",
		Handler: router,
	}
	log.Printf("Server started at %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
