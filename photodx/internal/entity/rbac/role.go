package rbac

type RoleID string

type Role struct {
	ID          RoleID
	Permissions []*Permission
}

// Predefined roles
var RoleGuest = Role{
	ID:          "Guest",
	Permissions: []*Permission{},
}
var RoleSuperUser = Role{
	ID: "SuperUser",
	Permissions: []*Permission{
		{Resource: ResourcePhotoStudio, Action: ActionGet},
		{Resource: ResourcePhotoStudio, Action: ActionList},
		{Resource: ResourcePhotoStudio, Action: ActionUpdate},
		{Resource: ResourcePhotoStudio, Action: ActionCreate},
		{Resource: ResourcePhotoStudio, Action: ActionUpdate},
		{Resource: ResourcePhotoStudio, Action: ActionDelete},
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
