package businesslogic

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

func (t *BusinessLogicImpl) CreateAccessToken(
	ctx context.Context,
	photoStudioMemberID entity.PhotoStudioMemberID,
) (string, error) {
	photoStudioMember, roles, photoStudio, err := t.Repository.GetPhotoStudioMember(
		ctx,
		photoStudioMemberID,
	)
	if err != nil {
		return "", terrors.Wrap(err)
	}
	if err != nil {
		return "", terrors.Wrap(err)
	}
	roleIDs := []rbac.RoleID{}
	for _, role := range roles {
		roleIDs = append(roleIDs, role.ID)
	}
	now := t.NowFunc()
	ttlMinutes := 5
	expiresAt := now.Add(time.Second * time.Duration(ttlMinutes) * 60)
	claims := auth.JWTClaimsAdminAccessToken{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   string(photoStudioMember.ID),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		Roles:         roleIDs,
		PhotoStudioID: photoStudio.ID,
	}
	tokenString, err := t.AccessTokenJWTCreator.CreateJWTToken(
		ctx,
		&claims,
	)
	if err != nil {
		return "", terrors.Wrap(err)
	}
	return tokenString, nil
}

func (t *BusinessLogicImpl) VerifyAccessToken(
	ctx context.Context,
	accessToken string,
) (entity.AdminPrincipal, error) {
	claims, err := t.AccessTokenJWTVerifier.VerifyJWTToken(ctx, accessToken, &auth.JWTClaimsAdminAccessToken{})
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
