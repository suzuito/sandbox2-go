package entity

type ClaimsRefreshToken interface {
	GetPrincipalID() PrincipalID
}

type ClaimsAccessToken interface {
	GetPrincipalID() PrincipalID
	GetRoles() []RoleID
}
