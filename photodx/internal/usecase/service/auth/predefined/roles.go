package predefined

import "github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"

var (
	RoleGuest = rbac.Role{
		ID:          "Guest",
		Permissions: []rbac.Permission{},
	}
	RoleSuperUser = rbac.Role{
		ID: "SuperUser",
		Permissions: []rbac.Permission{
			{Resource: "OrganizationBasicInfo", Action: rbac.ActionGet},
			{Resource: "OrganizationBasicInfo", Action: rbac.ActionList},
			{Resource: "OrganizationBasicInfo", Action: rbac.ActionUpdate},
			{Resource: "OrganizationBasicInfo", Action: rbac.ActionCreate},
			{Resource: "OrganizationBasicInfo", Action: rbac.ActionUpdate},
			{Resource: "OrganizationBasicInfo", Action: rbac.ActionDelete},
		},
	}
)

var AvailablePredefinedRoles = map[rbac.RoleID]*rbac.Role{}

func init() {
	roles := []rbac.Role{
		RoleGuest,
		RoleSuperUser,
	}
	for i := range roles {
		AvailablePredefinedRoles[roles[i].ID] = &roles[i]
	}
}
