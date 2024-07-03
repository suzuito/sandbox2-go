package entity

import (
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/pbrbac"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

type AdminPrincipal interface {
	GetPhotoStudioMemberID() PhotoStudioMemberID
	GetPhotoStudioID() PhotoStudioID
	GetRoles() []*rbac.Role
	GetPermissions() []*pbrbac.Permission
}

type AdminPrincipalImpl struct {
	PhotoStudioMemberID PhotoStudioMemberID
	PhotoStudioID       PhotoStudioID
	Roles               []*rbac.Role
}

func (t *AdminPrincipalImpl) GetPhotoStudioMemberID() PhotoStudioMemberID {
	return t.PhotoStudioMemberID
}

func (t *AdminPrincipalImpl) GetPhotoStudioID() PhotoStudioID {
	return t.PhotoStudioID
}

func (t *AdminPrincipalImpl) GetRoles() []*rbac.Role {
	return t.Roles
}

func (t *AdminPrincipalImpl) GetPermissions() []*pbrbac.Permission {
	permissions := []*pbrbac.Permission{}
	for _, role := range t.GetRoles() {
		permissions = append(permissions, role.Permissions...)
	}
	return permissions
}
