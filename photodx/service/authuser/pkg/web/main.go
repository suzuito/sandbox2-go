package web

import (
	"log/slog"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
	infra_repository "github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
	"gorm.io/gorm"
)

func Main(
	e *gin.Engine,
	l *slog.Logger,
	gormDB *gorm.DB,
	userRefreshTokenJWTCreator auth.JWTCreator,
	userRefreshTokenJWTVerifier auth.JWTVerifier,
	userAccessTokenJWTCreator auth.JWTCreator,
	userAccessTokenJWTVerifier auth.JWTVerifier,
) error {
	u := usecase.Impl{
		BusinessLogic: &businesslogic.Impl{
			Repository:                    &infra_repository.Impl{GormDB: gormDB, NowFunc: time.Now},
			NowFunc:                       time.Now,
			UserRefreshTokenJWTCreator:    userRefreshTokenJWTCreator,
			UserRefreshTokenJWTVerifier:   userRefreshTokenJWTVerifier,
			UserAccessTokenJWTCreator:     userAccessTokenJWTCreator,
			UserAccessTokenJWTVerifier:    userAccessTokenJWTVerifier,
			OAuth2LoginFlowStateGenerator: &proc.IDGeneratorImpl{},
			UserIDGenerator:               &proc.IDGeneratorImpl{},
		},
		CommonBusinessLogic: common_businesslogic.NewBusinessLogic(
			nil,
			userAccessTokenJWTVerifier,
		),
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
