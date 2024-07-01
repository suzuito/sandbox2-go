package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
)

func (t *Impl) Authorize(
	ctx context.Context,
	principal entity.Principal,
	requiredPermissions []*rbac.Permission,
) error {
	if len(requiredPermissions) <= 0 {
		return nil
	}
	if principal == nil {
		return &ForbiddenError{
			RequiredPermission: requiredPermissions[0],
		}
	}
	// TODO めっちゃ効率悪い。改善したい・・・
	// パーミションのマッチ
	roles := principal.GetRoles()
	for _, requiredPermission := range requiredPermissions {
		hasRequiredPermission := false
		for _, role := range roles {
			if hasRequiredPermission {
				break
			}
			for _, rolePermission := range role.Permissions {
				if requiredPermission.Action == rolePermission.Action && requiredPermission.Resource == rolePermission.Resource {
					// OK
					hasRequiredPermission = true
					break
				}
			}
		}
		if !hasRequiredPermission {
			return &ForbiddenError{
				RequiredPermission: requiredPermission,
			}
		}
	}
	return nil
}
