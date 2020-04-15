package data

import (
	"log"
	"time"
)

// User defines type user
type User struct {
	ID       int
	UUID     string
	FName    string
	LName    string
	Email    string
	Password string
	Created  time.Time
}

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

// Create method creates new user with provided data
func (user *User) Create() (err error) {
	stmt, err := DB.Prepare("INSERT INTO users (uuid, fname, lname, email, password, created_at) VALUES (?, ?, ?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Println("Prepare statement error", err)
		return
	}
	res, err := stmt.Exec(
		GenerateUUID(),
		user.FName,
		user.LName,
		user.Email,
		Encrypt(user.Password),
		time.Now(),
	)
	if err != nil {
		log.Println("Unable to insert data", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Unable to retrieve user", err)
	} else {
		log.Printf("User created with id:%d", id)
	}
	return
}

// CreateSession creates new session for existing user
func (user *User) CreateSession() (session Session, err error) {
	stmt, err := DB.Prepare("INSERT INTO sessions (uuid, fname, lname, email, usr_id, created_at) VALUES (?, ?, ?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Println("Prepare statement error", err)
		return
	}
	err = stmt.QueryRow(
		GenerateUUID(),
		user.FName,
		user.LName,
		user.Email,
		user.ID,
		time.Now(),
	).Scan(
		&session.ID,
		&session.UUID,
		&session.FName,
		&session.LName,
		&session.Email,
		&session.UserID,
		&session.Created,
	)
	return
}
