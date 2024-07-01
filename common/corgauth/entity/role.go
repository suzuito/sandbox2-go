package entity

type RoleID string

type Role struct {
	ID RoleID
}

func (t *Role) Validate() error {
	return nil
}

var MaxRolesPerPrincipal = 10
