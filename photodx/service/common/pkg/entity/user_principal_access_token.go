package entity

import (
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/pbrbac"
)

type UserPrincipalAccessToken interface {
	GetUserID() UserID
	GetPermissions() []*pbrbac.Permission
}
