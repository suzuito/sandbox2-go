package repository

import "fmt"

type EntryType string

const (
	EntryTypePhotoStudio       EntryType = "PhotoStudio"
	EntryTypePhotoStudioMember EntryType = "PhotoStudioMember"
)

type DuplicateEntryError struct {
	EntryType EntryType
	EntryID   string
}

func (t *DuplicateEntryError) Error() string {
	return fmt.Sprintf("duplicate entry %s.%s", t.EntryType, t.EntryID)
}

type NoEntryError struct {
	EntryType EntryType
	EntryID   string
}

func (t *NoEntryError) Error() string {
	return fmt.Sprintf("no entry %s.%s", t.EntryType, t.EntryID)
}
