package utils

import "time"

// Bod returns the beginning of the day (00:00:00) for the given time.
func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
