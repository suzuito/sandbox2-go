package serviceerror

import "fmt"

type NotFoundEntityError struct {
	EntityType string
	EntityID   string
}

func (t *NotFoundEntityError) Error() string {
	return fmt.Sprintf("not found %s : %s", t.EntityType, t.EntityID)
}

var PtrNotFoundEntityError *NotFoundEntityError

type AlreadyExistsEntityError struct {
	EntityType string
	EntityID   string
}

func (t *AlreadyExistsEntityError) Error() string {
	return fmt.Sprintf("already exists %s : %s", t.EntityType, t.EntityID)
}

var PtrAlreadyExistsEntityError *AlreadyExistsEntityError
