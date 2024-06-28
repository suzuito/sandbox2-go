package repository

import "fmt"

type NoEntryError struct {
	EntryType string
	EntryID   string
}

func (t *NoEntryError) Error() string {
	return fmt.Sprintf("no entry %s", t.EntryID)
}
