package repo

import (
	"fmt"
	"time"

	"github.com/LakshyaNegi/todos/entity"
)

func (r *repo) CreateTodo(task string) error {
	query :=
		fmt.Sprintf(`INSERT INTO todos (task, status, created_at, updated_at) 
		VALUES (?, "%s", datetime('now'), datetime('now'))`, entity.TodoStatusPending)

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = statement.Exec(task)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) CreateTodoWithDueDate(task string, dueDate time.Time) error {
	query :=
		fmt.Sprintf(`INSERT INTO todos (task, status, due_date, created_at, updated_at) 
		VALUES (?, "%s", ?, datetime('now'), datetime('now'))`, entity.TodoStatusPending)

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = statement.Exec(task, dueDate)
	if err != nil {
		return err
	}

	return nil
}
