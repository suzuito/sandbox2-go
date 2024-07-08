package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

func VerifyAdminAccessToken(
	ctx context.Context,
	verifier auth.JWTVerifier,
	accessToken string,
) (entity.AdminPrincipal, error) {
	claims, err := verifier.VerifyJWTToken(ctx, accessToken, &auth.JWTClaimsAdminAccessToken{})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	claimsAccessToken, ok := claims.(*auth.JWTClaimsAdminAccessToken)
	if !ok {
		return nil, terrors.Wrapf("cannot convert JWTClaims to JWTClaimsAccessToken")
	}
	principal := entity.AdminPrincipalImpl{
		PhotoStudioMemberID: entity.PhotoStudioMemberID(claimsAccessToken.Subject),
		PhotoStudioID:       claimsAccessToken.PhotoStudioID,
		Roles:               rbac.GetAvailablePredefinedRolesFromRoleID(claimsAccessToken.Roles),
	}
	return &principal, nil
}
