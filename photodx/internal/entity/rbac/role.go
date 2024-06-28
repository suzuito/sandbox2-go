package rbac

type RoleID string

type Role struct {
	ID          RoleID
	Permissions []Permission
}
