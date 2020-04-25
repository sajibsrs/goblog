// This file handles the application entry
// configuration and setup

package main

import (
	"github.com/julienschmidt/httprouter"
	"goblog/route"
	"goblog/route/user"
	"log"
	"net/http"
	"os"
)

func main() {
	router := httprouter.New()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to retrieve working directory")
	}
	router.ServeFiles("/static/*filepath", http.Dir(dir + "/static"))
	router.GET("/", route.Index)
	router.POST("/signup", user.New)
	router.GET("/signup", user.New)
	router.GET("/users", user.Index)
	router.GET("/users/view/:id", user.Index)
	router.GET("/users/new", user.New)
	router.POST("/users/create", user.New)
	server := &http.Server{
		Addr:    "127.0.0.1:2000",
		Handler: router,
	}
	log.Printf("Server started at %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
