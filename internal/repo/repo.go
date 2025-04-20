package repo

import (
	"database/sql"

	"github.com/LakshyaNegi/todos/db"
	"github.com/LakshyaNegi/todos/internal/entity"
)

type repo struct {
	db *sql.DB
}

var r *repo

func InitRepo() {
	r = &repo{
		db: db.Connect("data/todos.db"),
	}
}

func GetRepo() *repo {
	if r == nil {
		InitRepo()
	}

	return r
}

func (r *repo) runGetTodosQuery(query string, args ...any) ([]*entity.Todo, error) {
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []*entity.Todo{}

	for rows.Next() {
		todo := entity.Todo{}

		err := rows.Scan(&todo.ID, &todo.Task, &todo.Status, &todo.DueDate, &todo.CompletedAt, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, &todo)
	}

	return todos, nil
}
