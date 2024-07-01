package entity

import "github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"

type Principal interface {
	GetPhotoStudioMemberID() PhotoStudioMemberID
	GetPhotoStudioID() PhotoStudioID
	GetRoles() []*rbac.Role
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
