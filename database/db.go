package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDb() {
	var err error
	Db, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Verify the connection is actually working
	err = Db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)

	createTable()
}

func createTable() {

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL,
	password TEXT NOT NULL
)
	`
	_, err := Db.Exec(createUserTable)
	if err != nil {
		log.Fatal("Error creating user table: ", err)
	}

	// Fixed the typo from careateTable to createTableSQL
	createTableSQL := `CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(100) NOT NULL,
	description VARCHAR(255) NOT NULL,
	location VARCHAR(50) NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
	)`

	_, err = Db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating event table: ", err)
	}
}
