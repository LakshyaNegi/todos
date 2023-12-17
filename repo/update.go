package repo

import (
	"fmt"

	"github.com/LakshyaNegi/todos/entity"
)

func (r *repo) UpdateDoneByID(id int) error {
	query := fmt.Sprintf(
		`UPDATE todos SET completed_at = datetime('now'), status = %s, updated_at = datetime('now') WHERE id = ?`,
		entity.TodoStatusDone,
	)
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	row, err := statement.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("todo for id %v does not exist", id)
	}

	return nil
}

func (r *repo) UpdateTaskInProgressByID(id int) error {
	query := fmt.Sprintf(
		`UPDATE todos SET status = %s, updated_at = datetime('now') WHERE id = ?`,
		entity.TodoStatusInProgress,
	)

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	row, err := statement.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("todo for id %v does not exist", id)
	}

	return nil
}
