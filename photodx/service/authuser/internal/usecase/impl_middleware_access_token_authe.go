package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOMiddlewareAccessTokenAuthe struct {
	Principal common_entity.UserPrincipalAccessToken
}

func (t *Impl) MiddlewareAccessTokenAuthe(
	ctx context.Context,
	accessToken string,
) (*DTOMiddlewareAccessTokenAuthe, error) {
	principal, err := t.CommonBusinessLogic.VerifyUserAccessToken(ctx, accessToken)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOMiddlewareAccessTokenAuthe{
		Principal: principal,
	}, nil
}
