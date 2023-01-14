package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stevenfamy/go-web/config"
)

var DB *sql.DB

type dbConfig struct {
	dbHost     string
	dbPort     string
	dbName     string
	dbUser     string
	dbPassword string
}

func ConnectDatabase() {
	var config = dbConfig{
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_NAME"),
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
	}

	connectionString := config.dbUser + ":" + config.dbPassword + "@tcp(" + config.dbHost + ":" + config.dbPort + ")/" + config.dbName

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error())
	}

	DB = db
}
