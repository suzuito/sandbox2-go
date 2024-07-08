package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/entity"
)

type DTOMiddlewareRefreshTokenAuthe struct {
	Principal entity.AdminPrincipalRefreshToken
}

func (t *Impl) MiddlewareRefreshTokenAuthe(
	ctx context.Context,
	refreshToken string,
) (*DTOMiddlewareRefreshTokenAuthe, error) {
	principal, err := t.B.VerifyAdminRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOMiddlewareRefreshTokenAuthe{
		Principal: principal,
	}, nil
}
