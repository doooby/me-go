package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import sqlite3 driver
)

var DB *sql.DB

func InitDB(filepath string) {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		message TEXT,
		start_at TEXT NOT NULL,
		end_at TEXT,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL
	);`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}
