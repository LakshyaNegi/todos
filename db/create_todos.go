package db

import (
	"database/sql"
)

const (
	createTodosTable = `CREATE TABLE IF NOT EXISTS todos (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"task" TEXT NOT NULL,
		"status" TEXT NOT NULL,
		"due_date" TIMESTAMP,
		"completed_at" TIMESTAMP,
		"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		"updated_at" TIMESTAMP NOT NULL
	  );`
)

func createTable(db *sql.DB) error {
	statement, err := db.Prepare(createTodosTable)
	if err != nil {
		return err
	}

	statement.Exec()

	return nil
}
