package user

import (
	"goblog/database"
	"log"
	"regexp"
	"strings"
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

// Data displays user fields validation messages
type Data struct {
	*User
	Errors map[string]string
}

// Validate method checks for user fields validation and generates messages
// to display in the front end
func (msg *Data) Validate() map[string]string {
	var rxEmail = regexp.MustCompile(".+@.+\\..+")
	msg.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(msg.Email))
	if match == false {
		msg.Errors["email"] = "Please enter a valid email"
	}
	if strings.TrimSpace(msg.FName) == "" {
		msg.Errors["fname"] = "First name is required"
	}
	if strings.TrimSpace(msg.FName) == "" {
		msg.Errors["lname"] = "First name is required"
	}
	if len(msg.Password) < 6 {
		msg.Errors["pwd"] = "Password must be at least 6 characters"
	}
	return msg.Errors
}

// Create method creates new user with provided data
func (user *User) Create() (err error) {
	stmt, err := database.DB.Prepare("INSERT INTO users (uuid, fname, lname, email, password, created_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Prepare statement error", err)
		return
	}
	defer log.Fatal(stmt.Close())
	res, err := stmt.Exec(
		database.GenerateUUID(),
		user.FName,
		user.LName,
		user.Email,
		database.Encrypt(user.Password),
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
