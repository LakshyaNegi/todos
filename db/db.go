package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	db, err := sql.Open("sqlite3", "data/todos.db")
	if err != nil {
		log.Fatal("failed to open database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	err = createTable(db)
	if err != nil {
		log.Fatal("failed to create todos table:", err)
	}
}

func Connect(file string) *sql.DB {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal("failed to open database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	return db
}
