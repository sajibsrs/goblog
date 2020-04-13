// This file handles user routes of the application
// and serves templates accordingly

package handler

import (
	"goblog/data"
	"html/template"
	"net/http"
	"time"
)

// User handles request at /users/
func User(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"template/default/layout.html",
		"template/user/detail.html",
		"template/default/navigation.html"}

	tmpl := template.Must(template.ParseFiles(files...))
	_ = tmpl.ExecuteTemplate(w, "layout", r)
}

// CreateUser creates new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := data.User{
		UUID:     data.GenerateUUID(),
		FName:    "Sajidur",
		LName:    "Rahman",
		Email:    "sasjibsrs@gmail.com",
		Password: data.Encrypt("8080k"),
		Created:  time.Now(),
	}
	user.Create()
}
