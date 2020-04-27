package route

import (
	"github.com/julienschmidt/httprouter"
	"goblog/controller/user"
	"log"
	"net/http"
	"os"
)

// Route handles routing rules of the application
func Route(r *httprouter.Router) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to retrieve working directory")
	}
	r.ServeFiles("/static/*filepath", http.Dir(dir + "/static"))
	r.GET("/", user.Index)
	r.POST("/signup", user.New)
	r.GET("/signup", user.New)
	r.GET("/users", user.Index)
	r.GET("/users/view/:id", user.Index)
}
