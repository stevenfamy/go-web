package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	db, err := sql.Open("mysql", "root:rootlogistic@tcp(localhost:3306)/godev")

	if err != nil {
		panic(err.Error())
	}

	DB = db
}
