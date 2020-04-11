package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db, err := sql.Open("mysql", "root:8080k@tcp(localhost)/goblog")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Database connected")
}
