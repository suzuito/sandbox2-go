package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOMiddlewareAccessTokenAuthe struct {
	Principal entity.AdminPrincipalAccessToken
}

func (t *Impl) MiddlewareAccessTokenAuthe(
	ctx context.Context,
	accessToken string,
) (*DTOMiddlewareAccessTokenAuthe, error) {
	principal, err := t.B.VerifyAdminAccessToken(ctx, accessToken)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOMiddlewareAccessTokenAuthe{
		Principal: principal,
	}, nil
}
