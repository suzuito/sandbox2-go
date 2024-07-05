package web

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/auth/internal/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/inject"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

func Main(
	e *gin.Engine,
	l *slog.Logger,
	pool *sql.DB,
	adminRefreshTokenJWTPrivateKey string,
	adminAccessTokenJWTPrivateKey string,
	adminAccessTokenJWTPublicKey string,
) error {
	adminRefreshTokenJWTProcessor := auth.JWTHS{
		PrivateKey: []byte(adminRefreshTokenJWTPrivateKey),
	}
	adminAccessTokenJWTPrivateKeyBytes, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(adminAccessTokenJWTPrivateKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	adminAccessTokenJWTPublicKeyBytes, err := jwt.ParseRSAPublicKeyFromPEM([]byte(adminAccessTokenJWTPublicKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	b := businesslogic.Impl{
		Repository:                   inject.NewRepository(pool),
		SaltRepository:               inject.NewSaltRepository("foo"),
		PasswordHasher:               &proc.PasswordHasherMD5{},
		PhotoStudioMemberIDGenerator: &proc.IDGeneratorImpl{},
		PhotoStudioMemberInitialPasswordGenerator: &proc.InitialPasswordGeneratorImpl{},
		NowFunc:                      time.Now,
		AdminRefreshTokenJWTCreator:  &adminRefreshTokenJWTProcessor,
		AdminRefreshTokenJWTVerifier: &adminRefreshTokenJWTProcessor,
		AdminAccessTokenJWTCreator: &auth.JWTCreatorRS{
			PrivateKey: adminAccessTokenJWTPrivateKeyBytes,
		},
		AdminAccessTokenJWTVerifier: &auth.JWTVerifiers{
			Verifiers: []auth.JWTVerifier{
				&auth.JWTVerifierRS{
					PublicKey: adminAccessTokenJWTPublicKeyBytes,
				},
			},
		},
	}
	u := usecase.Impl{
		B: &b,
		L: l,
	}
	w := internal_web.Impl{
		U: &u,
		L: l,
		P: &presenter.Impl{},
	}
	// Authを担うAPI
	auth := e.Group("auth")
	auth.POST("login", w.AuthPostLogin)
	{
		x := auth.Group("x")
		x.Use(w.MiddlewareRefreshTokenAuthe)
		x.Use(w.MiddlewareRefreshTokenAutho)
		x.POST(
			"refresh",
			w.AuthPostRefresh,
		)
	}
	return nil
}
