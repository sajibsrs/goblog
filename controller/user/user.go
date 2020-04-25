// This file handles user routes of the application
// and serves templates accordingly

package user

import (
	"github.com/julienschmidt/httprouter"
	"goblog/data"
	"goblog/data/model"
	"goblog/helper"
	"log"
	"net/http"
	"time"
)

// Index handles user root request
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.URL.Path != "/users" {
		http.NotFound(w, r)
		return
	}
	tmp := []string{
		"user.detail",
		"default.layout",
		"default.navigation",
	}
	helper.ProcessTemplates(w, "layout", r, tmp...)
}

// New handles new create user request
func New(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.URL.Path != "/signup" {
		http.NotFound(w, r)
		return
	}
	//var prob []string
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
		//passOne := r.PostFormValue("pass_one")
		//if len(passOne) < 6 {
		//	prob = append(prob, "Password should be at least 6 characters")
		//}
		//if passTwo := r.PostFormValue("pass_two"); passOne != passTwo {
		//	prob = append(prob, "Password doesn't match")
		//}
		//if prob != nil {
		//	helper.ProcessTemplates(w, "layout", prob, tmp...)
		//	return
		//}
		user := model.User{
			UUID:     data.GenerateUUID(),
			FName:    r.PostFormValue("fname"),
			LName:    r.PostFormValue("lname"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("pass_one"),
			Created:  time.Now(),
		}

		msg := &model.Message{
			User:   user,
		}
		m := msg.Validate()
		if user.Password != r.PostFormValue("pass_two") {
			m["Password"] = "Password doesn't match"
		}
		if len(m) > 0 {
			helper.ProcessTemplates(w, "layout", msg, tmp...)
			return
		}

		if err := user.Create(); err != nil {
			log.Println("Cannot create user", err)
		}
		http.Redirect(w, r, "/", 302)
	}
	helper.ProcessTemplates(w, "layout", "", tmp...)
}
