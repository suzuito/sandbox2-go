package service

import (
	"fmt"

	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
)

var ErrPasswordMismatch = fmt.Errorf("password mismatch")

type ForbiddenError struct {
	RequiredPermission *rbac.Permission
}

func (t *ForbiddenError) Error() string {
	return fmt.Sprintf("permission %s.%s is required", t.RequiredPermission.Resource, t.RequiredPermission.Action)
}
