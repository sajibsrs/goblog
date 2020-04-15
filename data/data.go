package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// DB define database variable
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:8080k@tcp(localhost)/goblog")
	if err != nil {
		log.Fatalln("Unable establish database connection", err)
	}
	return
}

// GenerateUUID generates a new uuid
func GenerateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Println("UUID generation failed", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// Encrypt function encrypts string with SHA-1
func Encrypt(str string) (crypt string) {
	crypt = fmt.Sprintf("%x", sha1.Sum([]byte(str)))
	return
}
