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

// Create method creates new user with provided data
func (user *User) Create() (err error) {
	stmt, err := DB.Prepare("INSERT INTO users (uuid, fname, lname, email, password, created_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Prepare statement error", err)
		return
	}
	defer log.Fatal(stmt.Close())
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
