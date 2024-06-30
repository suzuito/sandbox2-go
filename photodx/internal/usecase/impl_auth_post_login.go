package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

type DTOAuthPostLogin struct {
	PhotoStudioMember *entity.PhotoStudioMember
	RefreshToken      string
}

func (t *Impl) AuthPostLogin(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
	password string,
) (*DTOAuthPostLogin, error) {
	if err := t.S.VerifyPhotoStudioMemberPassword(
		ctx,
		photoStudioID,
		email,
		password,
	); err != nil {
		return nil, terrors.Wrap(err)
	}
	photoStudioMember, err := t.S.GetPhotoStudioMemberByEmail(
		ctx,
		photoStudioID,
		email,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	refreshToken, err := t.S.CreateRefreshToken(ctx, photoStudioMember.ID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAuthPostLogin{
		PhotoStudioMember: photoStudioMember,
		RefreshToken:      refreshToken,
	}, nil
}
