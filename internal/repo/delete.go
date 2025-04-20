package repo

import "fmt"

func (r *repo) DeleteByID(id int) error {
	query := `DELETE FROM todos WHERE id = ?`

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
