package entity

import (
	"database/sql"
	"log"
	"time"
)

type Todo struct {
	ID          int          `json:"id" db:"id"`
	Task        string       `json:"task" db:"task"`
	Status      TodoStatus   `json:"status" db:"status"`
	DueDate     sql.NullTime `json:"due_date" db:"due_date"`
	CompletedAt sql.NullTime `json:"completed_at" db:"completed_at"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

func (t *Todo) Show() {
	log.Printf(
		"id: %v, task: %v, status: %v, due date: %v, completed at: %v\n",
		t.ID, t.Task, t.Status, t.DueDate, t.CompletedAt,
	)
}
