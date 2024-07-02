package service

import (
	"fmt"
)

var ErrPasswordMismatch = fmt.Errorf("password mismatch")

type ForbiddenError struct {
	// RequiredPermission *rbac.Permission
	Message string
}

func (t *ForbiddenError) Error() string {
	// return fmt.Sprintf("permission %s.%s is required", t.RequiredPermission.Resource, t.RequiredPermission.Action)
	return t.Message
}
