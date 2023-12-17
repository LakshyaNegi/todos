package repo

import (
	"time"

	"github.com/LakshyaNegi/todos/entity"
)

func (r *repo) GetTodoById(id int) (*entity.Todo, error) {
	query := `SELECT * FROM todos WHERE id = ?`

	row := r.db.QueryRow(query, id)

	todo := entity.Todo{}

	err := row.Scan(&todo.ID, &todo.Task, &todo.Status, &todo.DueDate, &todo.CompletedAt, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *repo) GetTodos() ([]*entity.Todo, error) {
	query := `SELECT * FROM todos ORDER BY created_at DESC`

	return r.runGetTodosQuery(query)
}

func (r *repo) GetIncompleteTodosOrderedByDueDateAsc() ([]*entity.Todo, error) {
	query := `SELECT * FROM todos WHERE completed_at IS NULL ORDER BY due_date ASC NULLS LAST`

	return r.runGetTodosQuery(query)
}

func (r *repo) GetCompletedTodos() ([]*entity.Todo, error) {
	query := `SELECT * FROM todos WHERE completed_at IS NOT NULL ORDER BY completed_at DESC`

	return r.runGetTodosQuery(query)
}

func (r *repo) GetCompletedTodosAfter(date time.Time) ([]*entity.Todo, error) {
	query := `SELECT * FROM todos WHERE completed_at > ? ORDER BY completed_at DESC`

	return r.runGetTodosQuery(query, date)
}

func (r *repo) GetDueTodos() ([]*entity.Todo, error) {
	query := `SELECT * FROM todos WHERE due_date IS NOT NULL AND completed_at IS NULL ORDER BY due_date ASC`

	return r.runGetTodosQuery(query)
}

func (r *repo) GetDueTodosBefore(date time.Time) ([]*entity.Todo, error) {
	query := `SELECT * FROM todos WHERE due_date < ? AND completed_at IS NULL ORDER BY due_date ASC`

	return r.runGetTodosQuery(query, date)
}

func (r *repo) GetPendingTodos() ([]*entity.Todo, error) {
	query := `SELECT * FROM todos WHERE completed_at IS NULL ORDER BY due_date ASC NULLS LAST`

	return r.runGetTodosQuery(query)
}

func (r *repo) GetOverdueTodos() ([]*entity.Todo, error) {
	query := `SELECT * FROM todos WHERE due_date < ? AND completed_at IS NULL ORDER BY due_date ASC`

	return r.runGetTodosQuery(query, time.Now())
}
