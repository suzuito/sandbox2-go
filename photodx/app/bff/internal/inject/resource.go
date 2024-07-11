package inject

import (
	"context"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/app/bff/internal/environment"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"gorm.io/gorm"
)

type Resource struct {
	GormDB *gorm.DB
	Logger *slog.Logger

	AdminAccessTokenJWTVerifier  auth.JWTVerifier
	AdminAccessTokenJWTCreator   auth.JWTCreator
	AdminRefreshTokenJWTVerifier auth.JWTVerifier
	AdminRefreshTokenJWTCreator  auth.JWTCreator
	UserAccessTokenJWTVerifier   auth.JWTVerifier
	UserAccessTokenJWTCreator    auth.JWTCreator
	UserRefreshTokenJWTVerifier  auth.JWTVerifier
	UserRefreshTokenJWTCreator   auth.JWTCreator
}

func (t *Resource) Close() {
	t.Logger.Info("Close resources")
}

func NewResource(
	ctx context.Context,
	env *environment.Environment,
) (*Resource, error) {
	resource := Resource{}
	switch env.Env {
	case "local":
		if err := setLocalResource(env, &resource); err != nil {
			return nil, terrors.Wrap(err)
		}
	case "prd":
		if err := setPrdResource(env, &resource); err != nil {
			return nil, terrors.Wrap(err)
		}
	default:
		return nil, terrors.Wrapf("undefined env resource : %s", env.Env)
	}
	if err := setJWTResource(env, &resource); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &resource, nil
}

func setJWTResource(
	env *environment.Environment,
	resource *Resource,
) error {
	// AdminAccessToken
	adminRefreshTokenProcessor := auth.JWTHS256{
		PrivateKey: []byte(env.JWTAdminRefreshTokenSigningPrivateKey),
	}
	resource.AdminRefreshTokenJWTCreator = &adminRefreshTokenProcessor
	resource.AdminRefreshTokenJWTVerifier = &adminRefreshTokenProcessor
	adminAccessTokenJWTPrivateKeyBytes, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(env.JWTAdminAccessTokenSigningPrivateKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	resource.AdminAccessTokenJWTCreator = &auth.JWTCreatorRS256{
		PrivateKey: adminAccessTokenJWTPrivateKeyBytes,
	}
	adminAccessTokenJWTPublicKeyBytes, err := jwt.ParseRSAPublicKeyFromPEM([]byte(env.JWTAdminAccessTokenSigningPublicKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	resource.AdminAccessTokenJWTVerifier = &auth.JWTVerifiers{
		Verifiers: []auth.JWTVerifier{
			&auth.JWTVerifierRS256{
				PublicKey: adminAccessTokenJWTPublicKeyBytes,
			},
		},
	}
	// UserAccessToken
	userRefreshTokenProcessor := auth.JWTHS256{
		PrivateKey: []byte(env.JWTUserRefreshTokenSigningPrivateKey),
	}
	resource.UserRefreshTokenJWTCreator = &userRefreshTokenProcessor
	resource.UserRefreshTokenJWTVerifier = &userRefreshTokenProcessor
	userAccessTokenJWTPrivateKeyBytes, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(env.JWTUserAccessTokenSigningPrivateKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	resource.UserAccessTokenJWTCreator = &auth.JWTCreatorRS256{
		PrivateKey: userAccessTokenJWTPrivateKeyBytes,
	}
	userAccessTokenJWTPublicKeyBytes, err := jwt.ParseRSAPublicKeyFromPEM([]byte(env.JWTUserAccessTokenSigningPublicKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	resource.UserAccessTokenJWTVerifier = &auth.JWTVerifiers{
		Verifiers: []auth.JWTVerifier{
			&auth.JWTVerifierRS256{
				PublicKey: userAccessTokenJWTPublicKeyBytes,
			},
		},
	}
	return nil
}
