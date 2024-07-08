package repository

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

type modelPhotoStudioMemberRole struct {
	PhotoStudioMemberID entity.PhotoStudioMemberID `gorm:"not null"`
	RoleID              rbac.RoleID                `gorm:"not null"`
	CreatedAt           time.Time                  `gorm:"not null"`
}

func (t *modelPhotoStudioMemberRole) TableName() string {
	return "photo_studio_member_roles"
}

type modelPhotoStudioMemberRoles []*modelPhotoStudioMemberRole

func (t *modelPhotoStudioMemberRoles) ToEntity() []*rbac.Role {
	roleIDs := []rbac.RoleID{}
	for _, role := range *t {
		roleIDs = append(roleIDs, role.RoleID)
	}
	return rbac.GetAvailablePredefinedRolesFromRoleID(roleIDs)
}
