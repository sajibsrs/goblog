// This file handles user routes of the application
// and serves templates accordingly

package user

import (
	"github.com/julienschmidt/httprouter"
	"goblog/database"
	"goblog/helper"
	"goblog/middleware/auth"
	"goblog/model"
	"log"
	"net/http"
	"time"
)

// Index handles user root request
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = auth.ValidateSession(w, r)
	users := All()
	tmp := []string{
		"user.list",
		"default.layout",
		"default.navigation",
	}
	helper.ProcessTemplates(w, "layout", users, tmp...)
}

// Index handles user root request
func View(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = auth.ValidateSession(w, r)
	users := All()
	tmp := []string{
		"user.list",
		"default.layout",
		"default.navigation",
	}
	helper.ProcessTemplates(w, "layout", users, tmp...)
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
		usr := model.User{
			UUID:     database.GenerateUUID(),
			FName:    r.PostFormValue("fname"),
			LName:    r.PostFormValue("lname"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("pass_one"),
			Created:  time.Now(),
		}
		data := model.Data{
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

// All returns all users from database
func All() []model.User {
	result, err := database.DB.Query(`SELECT id, uuid, fname, lname, email, created_at  FROM users`)
	if err != nil {
		log.Println("Get all session failed", err)
	}
	var users []model.User
	for result.Next() {
		var user model.User
		err := result.Scan(&user.ID, &user.UUID, &user.FName, &user.LName, &user.Email, &user.Created)
		if err != nil {
			log.Println("Unable to get user", err)
		} else {
			users = append(users, user)
		}
	}
	return users
}