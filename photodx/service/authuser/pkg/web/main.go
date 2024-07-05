package web

import (
	"database/sql"
	"log/slog"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/inject"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/oauth2loginflow"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

func Main(
	e *gin.Engine,
	l *slog.Logger,
	pool *sql.DB,
	userRefreshTokenJWTPrivateKey string,
	userAccessTokenJWTPrivateKey string,
	userAccessTokenJWTPublicKey string,
) error {
	userRefreshTokenJWTProcessor := auth.JWTHS{
		PrivateKey: []byte(userRefreshTokenJWTPrivateKey),
	}
	userAccessTokenJWTPrivateKeyBytes, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(userAccessTokenJWTPrivateKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	userAccessTokenJWTPublicKeyBytes, err := jwt.ParseRSAPublicKeyFromPEM([]byte(userAccessTokenJWTPublicKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	b := businesslogic.Impl{
		Repository:                  inject.NewRepository(pool),
		NowFunc:                     time.Now,
		UserRefreshTokenJWTCreator:  &userRefreshTokenJWTProcessor,
		UserRefreshTokenJWTVerifier: &userRefreshTokenJWTProcessor,
		UserAccessTokenJWTCreator: &auth.JWTCreatorRS{
			PrivateKey: userAccessTokenJWTPrivateKeyBytes,
		},
		UserAccessTokenJWTVerifier: &auth.JWTVerifiers{
			Verifiers: []auth.JWTVerifier{
				&auth.JWTVerifierRS{
					PublicKey: userAccessTokenJWTPublicKeyBytes,
				},
			},
		},
		OAuth2LoginFlowStateGenerator: &proc.IDGeneratorImpl{},
		UserIDGenerator:               &proc.IDGeneratorImpl{},
	}
	u := usecase.Impl{
		B: &b,
		L: l,
		OAuth2ProviderLINE: &oauth2loginflow.Provider{
			ClientID:     "2005761043",
			ClientSecret: "3250327d6ab0c0f92938d37e6ff87750",
		},
	}
	oauth2RedirectURL, _ := url.Parse("http://localhost:8082/authuser/x/callback")
	w := internal_web.Impl{
		U:                 &u,
		P:                 &presenter.Impl{},
		OAuth2RedirectURL: *oauth2RedirectURL,
	}
	authuser := e.Group("authuser")
	// authuser.POST("login", func(ctx *gin.Context) {}) // Password login
	{
		x := authuser.Group("x")
		x.GET("callback", w.GetCallback)
	}
	{
		authorize := authuser.Group("authorize")
		authorize.GET("line", w.GetAuthorizeLine)
	}
	return nil
}
