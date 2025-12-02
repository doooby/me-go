package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const DBPath = "var/db.sqlite"

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", DBPath)
	if err != nil {
		log.Fatal(err)
	}

	validateDatabase()
}

func validateDatabase() {
	var count int
	err := DB.QueryRow("SELECT COUNT(name) FROM sqlite_master WHERE type='table' AND name='tasks';").Scan(&count)
	if err != nil || count == 0 {
		log.Fatal("DB: table 'tasks' doesn't exist. Use `bin/create_db.sh` or import from backup manually.")
	}
}
