package model

import (
	"goblog/database"
	"log"
	"time"
)

// Session defines type session
type Session struct {
	ID      int
	UUID    string
	FName   string
	LName   string
	Email   string
	UserID  int
	Created time.Time
}

// Validate validates a session against given session UUID
func (session *Session) Validate() (valid bool, err error) {
	err = database.DB.QueryRow("SELECT id, uuid, fname, lname, email, usr_id, created_at FROM sessions WHERE uuid = ?", session.UUID).
		Scan(&session.ID, &session.UUID, &session.FName, &session.LName, &session.Email, &session.UserID, &session.Created)
	if err != nil {
		valid = false
		return
	}
	if session.ID != 0 {
		valid = true
	}
	return
}

// GetSessionByID returns a session based on given id
func GetSessionByID(id int64) (session Session, err error) {
	err = database.DB.QueryRow("SELECT id, uuid, fname, lname, email, usr_id, created_at FROM sessions WHERE id = ?", id).
		Scan(&session.ID, &session.UUID, &session.FName, &session.LName, &session.Email, &session.UserID, &session.Created)
	if err != nil {
		log.Println("Get session by id query failed", err)
		return
	}else {
		log.Println("Session retrieved by id successfully")
	}
	return
}
