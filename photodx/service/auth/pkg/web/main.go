package web

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/businesslogic"
	infra_repository "github.com/suzuito/sandbox2-go/photodx/service/auth/internal/infra/repository"
	infra_saltrepository "github.com/suzuito/sandbox2-go/photodx/service/auth/internal/infra/saltrepository"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/auth/internal/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
	"gorm.io/gorm"
)

func Main(
	e *gin.Engine,
	l *slog.Logger,
	gormDB *gorm.DB,
	adminRefreshTokenJWTPrivateKey string,
	adminAccessTokenJWTPrivateKey string,
	adminAccessTokenJWTPublicKey string,
) error {
	adminRefreshTokenJWTProcessor := auth.JWTHS256{
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
		Repository: &infra_repository.Impl{
			GormDB:  gormDB,
			NowFunc: time.Now,
		},
		SaltRepository:                            &infra_saltrepository.Impl{},
		PasswordHasher:                            &proc.PasswordHasherMD5{},
		PhotoStudioMemberIDGenerator:              &proc.IDGeneratorImpl{},
		PhotoStudioMemberInitialPasswordGenerator: &proc.InitialPasswordGeneratorImpl{},
		NowFunc:                      time.Now,
		AdminRefreshTokenJWTCreator:  &adminRefreshTokenJWTProcessor,
		AdminRefreshTokenJWTVerifier: &adminRefreshTokenJWTProcessor,
		AdminAccessTokenJWTCreator: &auth.JWTCreatorRS256{
			PrivateKey: adminAccessTokenJWTPrivateKeyBytes,
		},
		AdminAccessTokenJWTVerifier: &auth.JWTVerifiers{
			Verifiers: []auth.JWTVerifier{
				&auth.JWTVerifierRS256{
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
	// 後で消す
	auth.POST("super_init", w.SuperPostInit)
	return nil
}
