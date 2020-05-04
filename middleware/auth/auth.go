package auth

import (
	"github.com/julienschmidt/httprouter"
	"goblog/database"
	"goblog/helper"
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
	usr, _ := model.GetUserByEmail(r.PostFormValue("email"))
	if usr.Password == database.Encrypt(r.PostFormValue("password")) {
		sessionID, _ := usr.CreateSession()
		session, _ := GetSessionByID(sessionID)
		cookie := http.Cookie{
			Name: "goblog_session_cookie",
			Value: session.UUID,
			HttpOnly: true,
			Expires: time.Now().Add(24 * time.Hour),
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func CleanupSession(){
	result, err := database.DB.Query(`SELECT * FROM sessions`)
	if err != nil {
		log.Println("Get all session failed", err)
	}
	for result.Next() {
		var session Session
		err = result.Scan(&session.ID, &session.UUID, &session.FName, &session.LName, &session.Email, &session.UserID,
			&session.Created)
		if err != nil {
			log.Println("Unable to initialize sessions", err)
		} else {
			timeNow := time.Now()
			sessionTime := session.Created
			diff := timeNow.Sub(sessionTime).Hours()
			if diff > 24 {
				stmt, _ := database.DB.Prepare(`DELETE FROM sessions WHERE id = ?`)
				_, err := stmt.Exec(session.ID)
				if err != nil {
					log.Println("Unable to remove session", err)
				}
			}
		}
	}
}

// ValidateSession gets existing cookie
// and checks if that session exists
func ValidateSession(w http.ResponseWriter, r *http.Request) (session Session, err error) {
	cookie, err := r.Cookie("goblog_session_cookie")
	if err == nil {
		session = Session{UUID: cookie.Value}
		if valid, _ := session.Validate(); !valid {
			cookie := http.Cookie{
				Name: "goblog_session_cookie",
				Value: session.UUID,
				HttpOnly: true,
				Expires: time.Now(),
				MaxAge: -1,
			}
			http.SetCookie(w, &cookie)
			log.Println("Invalid session", err)
			http.Redirect(w, r, "/login", 302)
		}
	} else {
		log.Println("Session doesn't exists", err)
		http.Redirect(w, r, "/login", 302)
	}
	return
}

// GetSessionByID returns a session based on given id
func GetSessionByID(id int64) (session Session, err error) {
	err = database.DB.QueryRow(`SELECT id, uuid, fname, lname, email, usr_id, created_at FROM sessions 
		WHERE id = ?`, id).Scan(&session.ID, &session.UUID, &session.FName, &session.LName, &session.Email,
		&session.UserID, &session.Created)
	if err != nil {
		log.Println("Get session by id query failed", err)
		return
	}else {
		log.Println("Session retrieved by id successfully")
	}
	return
}

// Login handles existing user login
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	tmp := []string{
		"user.login",
		"default.layout",
		"default.navigation",
	}
	helper.ProcessTemplates(w, "layout", nil, tmp...)
}
