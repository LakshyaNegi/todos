package entity

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type TodoStatus string

const (
	TodoStatusPending    TodoStatus = "pending"
	TodoStatusDone       TodoStatus = "done"
	TodoStatusInProgress TodoStatus = "in_progress"
)

// String returns the string version of TodoStatus.
func (s *TodoStatus) String() string {
	return string(*s)
}

// Value implements the driver.Valuer interface.
func (s *TodoStatus) Value() (driver.Value, error) {
	if *s == "" {
		return nil, errors.New("unable to store TodoStatus: value cannot be empty")
	}
	return string(*s), nil
}

// Scan implements the sql.Scanner interface
func (s *TodoStatus) Scan(v interface{}) error {
	var finalValue string
	switch vv := v.(type) {
	case []byte:
		finalValue = string(vv)
	case string:
		finalValue = vv
	default:
		// The value is neither string or bytes. The value is stored in wrong format.
		return fmt.Errorf("unknown supported value type for TodoStatus: %T", v)
	}

	if finalValue == "" {
		return errors.New("unable to load TodoStatus: value is empty")
	}
	*s = TodoStatus(finalValue)

	return nil
}
