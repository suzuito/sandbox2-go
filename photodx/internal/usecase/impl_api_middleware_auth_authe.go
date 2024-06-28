package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

type DTOAPIMiddlewareAuthAuthe struct {
	Principal entity.Principal
}

func (t *Impl) APIMiddlewareAuthAuthe(
	ctx context.Context,
	accessToken string,
) (*DTOAPIMiddlewareAuthAuthe, error) {
	principal, err := t.S.VerifyAccessToken(ctx, accessToken)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIMiddlewareAuthAuthe{
		Principal: principal,
	}, nil
}
