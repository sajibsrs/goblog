package auth

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"goblog/database"
	"goblog/model"
	"log"
	"net/http"
	"time"
)

// CreateSession handles authentication of the user
func CreateSession(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Unable to parse form", err)
		return
	}
	expiration := time.Now().Add(24 * time.Hour)
	usr, _ := model.GetUserByEmail(r.PostFormValue("email"))
	if usr.Password == database.Encrypt(r.PostFormValue("password")) {
		sessionID, _ := usr.CreateSession()
		session, _ := GetSessionByID(sessionID)
		cookie := http.Cookie{
			Name: "goblog_session_cookie",
			Value: session.UUID,
			HttpOnly: true,
			Expires: expiration,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func ValidateSession(w http.ResponseWriter, r *http.Request) (session Session, err error) {
	cookie, err := r.Cookie("goblog_session_cookie")
	if err == nil {
		session = Session{UUID: cookie.Value}
		if valid, _ := session.Validate(); !valid {
			err = errors.New("invalid session")
		}
	}
	return
}