package entity

import (
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type AdminPrincipalRefreshToken interface {
	GetPhotoStudioMemberID() common_entity.PhotoStudioMemberID
}
