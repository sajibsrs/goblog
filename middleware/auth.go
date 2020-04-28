package middleware

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"goblog/database"
	"goblog/model"
	"log"
	"net/http"
)

// Auth handles authentication of the user
func Auth(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Unable to parse form", err)
		return
	}
	usr, _ := model.GetUserByEmail(r.PostFormValue("email"))
	if usr.Password == database.Encrypt(r.PostFormValue("password")) {
		sessionID, _ := usr.CreateSession()
		session, _ := model.GetSessionByID(sessionID)
		log.Println("session id", session.UUID)
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func ValidateAuth(w http.ResponseWriter, r *http.Request) (session model.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		session = model.Session{UUID: cookie.Value}
		if valid, _ := session.Validate(); !valid {
			err = errors.New("invalid session")
		}
	}
	return
}