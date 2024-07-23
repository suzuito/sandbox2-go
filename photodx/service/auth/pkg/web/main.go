package web

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/businesslogic"
	infra_repository "github.com/suzuito/sandbox2-go/photodx/service/auth/internal/infra/repository"
	infra_saltrepository "github.com/suzuito/sandbox2-go/photodx/service/auth/internal/infra/saltrepository"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/auth/internal/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
	"gorm.io/gorm"
)

func Main(
	e *gin.Engine,
	l *slog.Logger,
	gormDB *gorm.DB,
	adminRefreshTokenJWTCreator auth.JWTCreator,
	adminRefreshTokenJWTVerifier auth.JWTVerifier,
	adminAccessTokenJWTCreator auth.JWTCreator,
	adminAccessTokenJWTVerifier auth.JWTVerifier,
	webPushVAPIDPrivateKey string,
	webPushVAPIDPublicKey string,
) error {
	u := usecase.Impl{
		BusinessLogic: &businesslogic.Impl{
			Repository: &infra_repository.Impl{
				GormDB:  gormDB,
				NowFunc: time.Now,
			},
			SaltRepository:                            &infra_saltrepository.Impl{},
			PasswordHasher:                            &proc.PasswordHasherMD5{},
			PhotoStudioMemberIDGenerator:              &proc.IDGeneratorImpl{},
			PhotoStudioMemberInitialPasswordGenerator: &proc.InitialPasswordGeneratorImpl{},
			NowFunc:                      time.Now,
			AdminRefreshTokenJWTCreator:  adminRefreshTokenJWTCreator,
			AdminRefreshTokenJWTVerifier: adminRefreshTokenJWTVerifier,
			AdminAccessTokenJWTCreator:   adminAccessTokenJWTCreator,
			AdminAccessTokenJWTVerifier:  adminAccessTokenJWTVerifier,
			WebPushVAPIDPrivateKey:       webPushVAPIDPrivateKey,
			WebPushVAPIDPublicKey:        webPushVAPIDPublicKey,
		},
		CommonBusinessLogic: common_businesslogic.NewBusinessLogic(
			adminAccessTokenJWTVerifier,
			nil,
		),
		L: l,
	}
	p := presenter.Impl{}
	w := internal_web.Impl{
		U: &u,
		L: l,
		P: &p,
	}
	res := func(ctx *gin.Context, r any, err error) {
		common_web.Response(
			ctx,
			l,
			&p,
			r,
			err,
			&common_web.DefaultWebResponseOption,
		)
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
	{
		y := auth.Group("y")
		y.Use(w.MiddlewareAccessTokenAuthe)
		y.GET(
			"init",
			w.MiddlewareAccessTokenAuthe,
			common_web.MiddlewareAdminAccessTokenAutho(
				l,
				`permissions.exists(
    				p,
			        p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "read".matches(p.action)
		        ) &&
				permissions.exists(
    				p,
			        p.resource == "PhotoStudioMember" && adminPrincipalPhotoStudioMemberId.matches(p.target) && "read".matches(p.action)
		        )
				`,
				w.P,
			),
			w.AuthGetInit,
		)
		{
			pushAPI := y.Group("push_api")
			pushAPI.PUT(
				"push_subscription",
				func(ctx *gin.Context) {
					subscription := webpush.Subscription{}
					if err := ctx.BindJSON(&subscription); err != nil {
						p.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
							Message: err.Error(),
						})
						return
					}
					dto, err := u.AuthPutPushSubscription(ctx, common_web.CtxGetAdminPrincipalAccessToken(ctx), &subscription)
					res(ctx, dto, err)
				},
			)
		}
	}
	return nil
}
