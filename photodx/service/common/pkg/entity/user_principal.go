package entity

import (
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/pbrbac"
)

type UserPrincipal interface {
	GetUserID() UserID
	GetPermissions() []*pbrbac.Permission
}
