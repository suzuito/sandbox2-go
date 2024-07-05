package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) VerifyUserAccessToken(
	ctx context.Context,
	accessToken string,
) (entity.UserPrincipal, error) {
	claims, err := t.UserAccessTokenJWTVerifier.VerifyJWTToken(ctx, accessToken, nil)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	claimsAccessToken, ok := claims.(*auth.JWTClaimsUserAccessToken)
	if !ok {
		return nil, terrors.Wrapf("cannot convert JWTClaims to JWTClaimsUserAccessToken")
	}
	return claimsAccessToken, nil
}
