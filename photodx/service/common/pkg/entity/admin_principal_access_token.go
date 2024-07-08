package entity

import (
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/pbrbac"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

type AdminPrincipalAccessToken interface {
	GetPhotoStudioMemberID() PhotoStudioMemberID
	GetPhotoStudioID() PhotoStudioID
	GetRoles() []*rbac.Role
	GetPermissions() []*pbrbac.Permission
}

type AdminPrincipalAccessTokenImpl struct {
	PhotoStudioMemberID PhotoStudioMemberID
	PhotoStudioID       PhotoStudioID
	Roles               []*rbac.Role
}

func (t *AdminPrincipalAccessTokenImpl) GetPhotoStudioMemberID() PhotoStudioMemberID {
	return t.PhotoStudioMemberID
}

func (t *AdminPrincipalAccessTokenImpl) GetPhotoStudioID() PhotoStudioID {
	return t.PhotoStudioID
}

func (t *AdminPrincipalAccessTokenImpl) GetRoles() []*rbac.Role {
	return t.Roles
}

func (t *AdminPrincipalAccessTokenImpl) GetPermissions() []*pbrbac.Permission {
	permissions := []*pbrbac.Permission{}
	for _, role := range t.GetRoles() {
		permissions = append(permissions, role.Permissions...)
	}
	return permissions
}
