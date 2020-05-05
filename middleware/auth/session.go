package auth

import (
	"goblog/database"
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
	err = database.DB.QueryRow(`SELECT id, uuid, fname, lname, email, usr_id, created_at FROM sessions 
		WHERE uuid = ?`, session.UUID).Scan(&session.ID, &session.UUID, &session.FName, &session.LName,
		&session.Email, &session.UserID, &session.Created)
	if err != nil {
		valid = false
		return
	}
	if session.ID != 0 {
		valid = true
	}
	return
}
