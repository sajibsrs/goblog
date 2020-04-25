package data

import (
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

// CreateSession creates new session for existing user
func (user *User) CreateSession() (session Session, err error) {
	stmt, err := DB.Prepare("INSERT INTO sessions (uuid, fname, lname, email, usr_id, created_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Prepare statement error", err)
		return
	}
	defer log.Fatal(stmt.Close())
	_, err = stmt.Exec(
		GenerateUUID(),
		user.FName,
		user.LName,
		user.Email,
		user.ID,
		time.Now(),
	)
	if err != nil {
		log.Println("Unable to create session data", err)
	}
	if err != nil {
		log.Println("Unable to retrieve session data", err)
	} else {
		log.Printf("User created")
	}
	return
}
