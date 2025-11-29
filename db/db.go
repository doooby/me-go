package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import sqlite3 driver
)

var DBPath = "var/db.sqlite"
var DB *sql.DB


func InitDB(filepath string) {
	var err error
	DB, err = sql.Open(DBPath, filepath)
	if err != nil {
		log.Fatal(err)
	}
	validateDatabase()

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

func validateDatabase() {
	var count int
	err := DB.QueryRow("SELECT COUNT(name) FROM sqlite_master WHERE type='table' AND name='tasks';").Scan(&count)
	if err != nil || count == 0 {
		log.Fatal("DB: table 'tasks' doesn't exist. Use `bin/create_db.sh` or import from backup manually.")
	}
}
