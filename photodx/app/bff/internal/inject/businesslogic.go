package inject

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/app/bff/internal/environment"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	common_inject "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/inject"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
)

func NewBusinessLogic(
	env *environment.Environment,
	logger *slog.Logger,
) (businesslogic.BusinessLogic, error) {
	mysqlConfig := mysql.Config{
		DBName:               env.DBName,
		User:                 env.DBUser,
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3308",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	pool, err := sql.Open(
		"mysql",
		mysqlConfig.FormatDSN(),
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	repo := common_inject.NewRepository(
		pool,
	)
	saltRepo := common_inject.NewSaltRepository("foo")
	jwths := auth.JWTHS{
		PrivateKey: []byte(env.JWTAdminRefreshTokenSigningPrivateKey),
	}
	adminAccessTokenJWTPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(env.JWTAdminAccessTokenSigningPrivateKey))
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	adminAccessTokenJWTPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(env.JWTAdminAccessTokenSigningPublicKey))
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	auth0Verifier, err := auth.NewJWTVerifierAuth0(env.Auth0Domain, env.Auth0Audience)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	b := common_inject.NewBusinessLogic(
		&proc.IDGeneratorImpl{},
		&proc.InitialPasswordGeneratorImpl{},
		&proc.PasswordHasherMD5{},
		repo,
		saltRepo,
		&jwths,
		&jwths,
		&auth.JWTCreatorRS{
			PrivateKey: adminAccessTokenJWTPrivateKey,
		},
		&auth.JWTVerifiers{
			Verifiers: []auth.JWTVerifier{
				&auth.JWTVerifierRS{
					PublicKey: adminAccessTokenJWTPublicKey,
				},
			},
		},
		&auth.JWTVerifiers{
			Verifiers: []auth.JWTVerifier{
				auth0Verifier,
			},
		},
		time.Now,
	)
	return b, nil
}
