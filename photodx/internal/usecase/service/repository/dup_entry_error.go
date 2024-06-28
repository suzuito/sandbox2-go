package repository

import "fmt"

type DuplicateEntryError struct {
	EntryType string
	EntryID   string
}

func (t *DuplicateEntryError) Error() string {
	return fmt.Sprintf("duplicate entry %s.%s", t.EntryType, t.EntryID)
}
