// This file handles user routes of the application
// and serves templates accordingly

package user

import (
	"github.com/julienschmidt/httprouter"
	"goblog/database"
	"goblog/helper"
	"goblog/model/user"
	"log"
	"net/http"
	"time"
)

// Index handles user root request
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmp := []string{
		"user.detail",
		"default.layout",
		"default.navigation",
	}
	helper.ProcessTemplates(w, "layout", r, tmp...)
}

// New handles new create user request
func New(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmp := []string{
		"user.new",
		"default.layout",
		"default.navigation",
	}
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Unable to parse form", err)
			return
		}
		usr := &user.User{
			UUID:     database.GenerateUUID(),
			FName:    r.PostFormValue("fname"),
			LName:    r.PostFormValue("lname"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("pass_one"),
			Created:  time.Now(),
		}
		data := user.Data{
			User:   usr,
		}
		msg := data.Validate()
		if usr.Password != r.PostFormValue("pass_two") {
			msg["pwd_match"] = "Password doesn't match"
		}
		if len(msg) > 0 {
			helper.ProcessTemplates(w, "layout", data, tmp...)
			return
		}
		if err := usr.Create(); err != nil {
			log.Println("Cannot create user", err)
		}
		http.Redirect(w, r, "/", 302)
	}
	helper.ProcessTemplates(w, "layout", nil, tmp...)
}
