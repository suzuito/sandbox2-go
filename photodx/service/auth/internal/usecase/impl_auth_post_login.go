package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAuthPostLogin struct {
	RefreshToken string
}

func (t *Impl) AuthPostLogin(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
	password string,
) (*DTOAuthPostLogin, error) {
	photoStudioMember, _, _, err := t.BusinessLogic.VerifyPhotoStudioMemberPassword(
		ctx,
		photoStudioID,
		email,
		password,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	refreshToken, err := t.BusinessLogic.CreateAdminRefreshToken(ctx, photoStudioMember.ID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAuthPostLogin{
		RefreshToken: refreshToken,
	}, nil
}
