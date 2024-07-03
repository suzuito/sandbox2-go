package entity

import (
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/pbrbac"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

type Principal interface {
	GetPhotoStudioMemberID() PhotoStudioMemberID
	GetPhotoStudioID() PhotoStudioID
	GetRoles() []*rbac.Role
	GetPermissions() []*pbrbac.Permission
}

type PrincipalImpl struct {
	PhotoStudioMemberID PhotoStudioMemberID
	PhotoStudioID       PhotoStudioID
	Roles               []*rbac.Role
}

func (t *PrincipalImpl) GetPhotoStudioMemberID() PhotoStudioMemberID {
	return t.PhotoStudioMemberID
}

func (t *PrincipalImpl) GetPhotoStudioID() PhotoStudioID {
	return t.PhotoStudioID
}

func (t *PrincipalImpl) GetRoles() []*rbac.Role {
	return t.Roles
}

func (t *PrincipalImpl) GetPermissions() []*pbrbac.Permission {
	permissions := []*pbrbac.Permission{}
	for _, role := range t.GetRoles() {
		permissions = append(permissions, role.Permissions...)
	}
	return permissions
}
