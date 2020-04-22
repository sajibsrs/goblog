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
	if err != nil {
		log.Println("Prepare statement error", err)
		return
	}
	defer stmt.Close()
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
	if err != nil {
		log.Println("Prepare statement error", err)
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(
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
