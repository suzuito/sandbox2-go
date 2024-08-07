package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity"
)

type DTOPostRefreshAccessToken struct {
	AccessToken string `json:"accessToken"`
}

func (t *Impl) PostRefreshAccessToken(
	ctx context.Context,
	principal entity.UserPrincipalRefreshToken,
) (*DTOPostRefreshAccessToken, error) {
	accessToken, err := t.BusinessLogic.CreateUserAccessToken(ctx, principal.GetUserID())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOPostRefreshAccessToken{AccessToken: accessToken}, nil
}
