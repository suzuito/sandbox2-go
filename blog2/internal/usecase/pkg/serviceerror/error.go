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
