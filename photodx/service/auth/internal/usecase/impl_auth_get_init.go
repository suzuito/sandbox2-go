package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAuthGetInit struct {
	PhotoStudio       *entity.PhotoStudio       `json:"photoStudio"`
	PhotoStudioMember *entity.PhotoStudioMember `json:"photoStudioMember"`
	VAPIDPublicKey    string                    `json:"vapidPublicKey"`
}

func (t *Impl) AuthGetInit(
	ctx context.Context,
	principal entity.AdminPrincipalAccessToken,
) (*DTOAuthGetInit, error) {
	member, _, photoStudio, err := t.BusinessLogic.GetPhotoStudioMember(ctx, principal.GetPhotoStudioMemberID())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	key, err := t.BusinessLogic.GetWebPushVAPIDPublicKey(ctx, principal.GetPhotoStudioMemberID())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAuthGetInit{
		PhotoStudio:       photoStudio,
		PhotoStudioMember: member,
		VAPIDPublicKey:    key,
	}, nil
}
