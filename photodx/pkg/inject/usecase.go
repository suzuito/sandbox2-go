package inject

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/photodx/internal/infra/saltrepository"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/auth"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/proc"
	"github.com/suzuito/sandbox2-go/photodx/pkg/environment"
)

func NewUsecase(
	env *environment.Environment,
	logger *slog.Logger,
) (usecase.Usecase, error) {
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
	fmt.Println(mysqlConfig.FormatDSN())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	repo := repository.Impl{
		Pool: pool,
	}
	saltRepo := saltrepository.Impl{
		Version: "foo",
	}
	jwths := auth.JWTHS{
		PrivateKey: []byte(env.JWTRefreshTokenSigningPrivateKey),
	}
	accessTokenJWTPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(env.JWTAccessTokenSigningPrivateKey))
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	accessTokenJWTPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(env.JWTAccessTokenSigningPublicKey))
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	s := service.Impl{
		PhotoStudioMemberIDGenerator:              &proc.IDGeneratorImpl{},
		PhotoStudioMemberInitialPasswordGenerator: &proc.InitialPasswordGeneratorImpl{},
		PasswordHasher:                            &proc.PasswordHasherMD5{},
		Repository:                                &repo,
		SaltRepository:                            &saltRepo,
		RefreshTokenJWTCreator:                    &jwths,
		RefreshTokenJWTVerifier:                   &jwths,
		AccessTokenJWTCreator: &auth.JWTCreatorRS{
			PrivateKey: accessTokenJWTPrivateKey,
		},
		AccessTokenJWTVerifier: &auth.JWTVerifiers{
			Verifiers: []auth.JWTVerifier{
				&auth.JWTVerifierRS{
					PublicKey: accessTokenJWTPublicKey,
				},
			},
		},
		NowFunc: time.Now,
	}
	u := usecase.Impl{
		S: &s,
		L: logger,
	}
	return &u, nil
}
