package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
)

func (t *Impl) MiddlewareAccessTokenAutho(
	ctx context.Context,
	principal entity.Principal,
	requiredPermissions []*rbac.Permission,
) error {
	return terrors.Wrap(t.S.Authorize(ctx, principal, requiredPermissions))
}
