package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOAPIPostGuest struct {
	RefreshToken string `json:"refreshToken"`
}

func (t *Impl) APIPostGuest(ctx context.Context) (*DTOAPIPostGuest, error) {
	user, err := t.BusinessLogic.CreateGuestUser(ctx)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	refreshToken, err := t.BusinessLogic.CreateUserRefreshToken(ctx, user.ID, true)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostGuest{
		RefreshToken: refreshToken,
	}, nil
}
