package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOAPIPostLogin struct {
	RefreshToken string `json:"refreshToken"`
}

func (t *Impl) APIPostLogin(ctx context.Context, email string, password string) (*DTOAPIPostLogin, error) {
	user, err := t.BusinessLogic.VerifyUserPassword(ctx, email, password)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	refreshToken, err := t.BusinessLogic.CreateUserRefreshToken(ctx, user.ID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostLogin{RefreshToken: refreshToken}, nil
}
