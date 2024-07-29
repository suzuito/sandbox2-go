package entity

import (
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type UserPrincipalRefreshToken interface {
	GetUserID() common_entity.UserID
	IsGuestUser() bool
}
