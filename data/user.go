package data

import (
	"fmt"
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
	fmt.Println(stmt)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		GenerateUUID(),
		user.FName,
		user.LName,
		user.Email,
		Encrypt(user.Password),
		time.Now(),
	).Scan(
		&user.ID,
		&user.UUID,
		&user.FName,
		&user.LName,
		&user.Email,
		&user.Created,
	)
	return
}

// CreateSession creates new session for existing user
func (user *User) CreateSession() (session Session, err error) {
	stmt, err := DB.Prepare("INSERT INTO sessions (uuid, fname, lname, email, usr_id, created_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
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
