package rbac

import "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/pbrbac"

type RoleID string

type Role struct {
	ID          RoleID
	Permissions []*pbrbac.Permission
}

// Predefined roles
var RoleGuest = Role{
	ID:          "Guest",
	Permissions: []*pbrbac.Permission{},
}
var RoleSuperUser = Role{
	ID: "SuperUser",
	Permissions: []*pbrbac.Permission{
		{Resource: "PhotoStudio", Target: ".*", Action: "create"},
		{Resource: "PhotoStudio", Target: ".*", Action: "read"},
		{Resource: "PhotoStudio", Target: ".*", Action: "update"},
		{Resource: "PhotoStudio", Target: ".*", Action: "delete"},
		{Resource: "PhotoStudioMember", Target: ".*", Action: "create"},
		{Resource: "PhotoStudioMember", Target: ".*", Action: "read"},
		{Resource: "PhotoStudioMember", Target: ".*", Action: "update"},
		{Resource: "PhotoStudioMember", Target: ".*", Action: "delete"},
	},
}

var AvailablePredefinedRoles = map[RoleID]*Role{}

func GetAvailablePredefinedRolesFromRoleID(
	roleIDs []RoleID,
) []*Role {
	roles := []*Role{}
	for _, roleID := range roleIDs {
		role, exists := AvailablePredefinedRoles[roleID]
		if !exists {
			continue
		}
		roles = append(roles, role)
	}
	return roles
}

func init() {
	roles := []Role{
		RoleGuest,
		RoleSuperUser,
	}
	for i := range roles {
		AvailablePredefinedRoles[roles[i].ID] = &roles[i]
	}
}
