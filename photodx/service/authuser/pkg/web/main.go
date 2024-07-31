package web

import (
	"log/slog"
	"net/http"
	"net/url"
	"time"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
	infra_gmailsmtp "github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/infra/gateway/mail/gmailsmtp"
	infra_repository "github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/web"
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
	passwordSalt string,
	userRefreshTokenJWTCreator auth.JWTCreator,
	userRefreshTokenJWTVerifier auth.JWTVerifier,
	userAccessTokenJWTCreator auth.JWTCreator,
	userAccessTokenJWTVerifier auth.JWTVerifier,
	oauth2ProviderLINEClientID string,
	oauth2ProviderLINEClientSecret string,
	oauth2ProviderLINERedirectURL string,
	frontURLString string,
	webPushVAPIDPrivateKey string,
	webPushVAPIDPublicKey string,
	userMailSenderGmailSmtpAccount string,
	userMailSenderGmailSmtpPassword string,
	userMailSenderGmailSmtpFromEmail string,
	userMailSenderGmailSmtpFromName string,
) error {
	b := businesslogic.Impl{
		Repository: &infra_repository.Impl{GormDB: gormDB, NowFunc: time.Now},
		UserMailSender: &infra_gmailsmtp.UserMailSender{
			Host:      "smtp.gmail.com",
			Port:      587,
			Account:   userMailSenderGmailSmtpAccount,
			Password:  userMailSenderGmailSmtpPassword,
			FromName:  userMailSenderGmailSmtpFromName,
			FromEmail: userMailSenderGmailSmtpFromEmail,
		},
		NowFunc:                                   time.Now,
		PasswordSalt:                              passwordSalt,
		PasswordHasher:                            &proc.PasswordHasherMD5{},
		UserRefreshTokenJWTCreator:                userRefreshTokenJWTCreator,
		UserRefreshTokenJWTVerifier:               userRefreshTokenJWTVerifier,
		UserAccessTokenJWTCreator:                 userAccessTokenJWTCreator,
		UserAccessTokenJWTVerifier:                userAccessTokenJWTVerifier,
		OAuth2LoginFlowStateGenerator:             &proc.IDGeneratorImpl{},
		UserIDGenerator:                           &proc.IDGeneratorImpl{},
		PromoteGuestUserConfirmationCodeGenerator: &proc.IDGeneratorImpl{},
		WebPushVAPIDPrivateKey:                    webPushVAPIDPrivateKey,
		WebPushVAPIDPublicKey:                     webPushVAPIDPublicKey,
	}
	u := usecase.Impl{
		BusinessLogic: &b,
		CommonBusinessLogic: common_businesslogic.NewBusinessLogic(
			nil,
			userAccessTokenJWTVerifier,
		),
		L: l,
		OAuth2ProviderLINE: &oauth2loginflow.Provider{
			ClientID:     oauth2ProviderLINEClientID,
			ClientSecret: oauth2ProviderLINEClientSecret,
		},
	}
	oauth2RedirectURL, err := url.Parse(oauth2ProviderLINERedirectURL)
	if err != nil {
		return terrors.Wrap(err)
	}
	frontURL, err := url.Parse(frontURLString)
	if err != nil {
		return terrors.Wrap(err)
	}
	p := presenter.Impl{}
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
	w := internal_web.Impl{
		L:                 l,
		U:                 &u,
		P:                 &p,
		OAuth2RedirectURL: *oauth2RedirectURL,
		FrontBaseURL:      *frontURL,
	}
	authuser := e.Group("authuser")
	// authuser.POST("login", func(ctx *gin.Context) {}) // Password login
	// ==== Debug START ====
	// authuser.GET("a", func(ctx *gin.Context) {
	// 	err := b.PushNotification(
	// 		ctx,
	// 		l,
	// 		entity.UserID(ctx.Query("userId")),
	// 		&entity.Notification{
	// 			ID: fmt.Sprintf("test-%d", time.Now().Unix()),
	// 			Type: entity.NotificationTypeChatMessage,
	// 			ChatMessageWrapper: &entity.ChatMessageWrapper{
	// 				ChatMessage: entity.ChatMessage{},
	//
	// 			},
	// 		},
	// 	)
	// 	if err != nil {
	// 		l.Error("", "err", err)
	// 	}
	// })
	// ==== Debug END ====
	authuser.POST("guest", func(ctx *gin.Context) {
		dto, err := u.APIPostGuest(ctx)
		res(ctx, dto, err)
	})
	{
		x := authuser.Group("x")
		x.GET("callback", w.GetCallback)
	}
	{
		authorize := authuser.Group("authorize")
		authorize.GET("line", w.GetAuthorizeLine)
	}
	{
		y := authuser.Group("y")
		y.Use(func(ctx *gin.Context) {
			refreshToken := common_web.ExtractBearerToken(ctx)
			if refreshToken == "" {
				ctx.Next()
				return
			}
			dto, err := u.MiddlewareRefreshTokenAuthe(ctx, refreshToken)
			if err != nil {
				l.Warn("refreshToken authe is failed", "err", err)
				ctx.Next()
				return
			}
			internal_web.CtxSetUserPrincipalRefreshToken(ctx, dto.Principal)
			ctx.Next()
		})
		y.POST(
			"refresh",
			func(ctx *gin.Context) {
				if internal_web.CtxGetUserPrincipalRefreshToken(ctx) == nil {
					p.JSON(ctx, http.StatusForbidden, common_web.ResponseError{
						Message: "unauthorized",
					})
					ctx.Abort()
					return
				}
				ctx.Next()
			},
			func(ctx *gin.Context) {
				dto, err := u.PostRefreshAccessToken(ctx, internal_web.CtxGetUserPrincipalRefreshToken(ctx))
				res(ctx, dto, err)
			},
		)
	}
	{
		z := authuser.Group("z")
		z.Use(func(ctx *gin.Context) {
			accessToken := common_web.ExtractBearerToken(ctx)
			if accessToken == "" {
				ctx.Next()
				return
			}
			dto, err := u.MiddlewareAccessTokenAuthe(ctx, accessToken)
			if err != nil {
				l.Warn("accessToken authe is failed", "err", err)
				ctx.Next()
				return
			}
			common_web.CtxSetUserPrincipalAccessToken(ctx, dto.Principal)
			ctx.Next()
		})
		z.GET(
			"init",
			// common_web.MiddlewareUserAccessTokenAutho(
			// 	l,
			// 	`
			// 		permissions.exists(
			// 			p,
			// 			p.resource == "PhotoStudio" && userPrincipalUserId.matches(p.target) && "read".matches(p.action)
			// 		)
			// 		`,
			// 	&p,
			// ),
			func(ctx *gin.Context) {
				dto, err := u.APIGetInit(ctx, common_web.CtxGetUserPrincipalAccessToken(ctx))
				res(ctx, dto, err)
			},
		)
		{
			pushAPI := z.Group("push_api")
			pushAPI.PUT(
				"push_subscription",
				// common_web.MiddlewareUserAccessTokenAutho(
				// 	l,
				// 	`
				// 		permissions.exists(
				// 			p,
				// 			p.resource == "PhotoStudio" && userPrincipalUserId.matches(p.target) && "read".matches(p.action)
				// 		)
				// 		`,
				// 	&p,
				// ),
				func(ctx *gin.Context) {
					subscription := webpush.Subscription{}
					if err := ctx.BindJSON(&subscription); err != nil {
						p.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
							Message: err.Error(),
						})
						return
					}
					dto, err := u.APIPutPushSubscription(ctx, common_web.CtxGetUserPrincipalAccessToken(ctx), &subscription)
					res(ctx, dto, err)
				},
			)
		}
		{
			promote := z.Group("promote")
			promote.POST("request", func(ctx *gin.Context) {
				body := struct {
					Email string `json:"email"`
				}{}
				if err := ctx.BindJSON(&body); err != nil {
					p.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
						Message: err.Error(),
					})
					return
				}
				dto, err := u.APIPostRequestPromoteGuestUser(
					ctx,
					*frontURL,
					common_web.CtxGetUserPrincipalAccessToken(ctx).GetUserID(),
					body.Email,
				)
				res(ctx, dto, err)
			})
			promote.POST("approve", func(ctx *gin.Context) {
				body := struct {
					Code     string `form:"code"`
					Password string `form:"password"`
				}{}
				if err := ctx.BindJSON(&body); err != nil {
					p.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
						Message: err.Error(),
					})
					return
				}
				dto, err := u.APIPostApprovePromoteGuestUser(
					ctx,
					common_web.CtxGetUserPrincipalAccessToken(ctx),
					body.Password,
					body.Code,
				)
				res(ctx, dto, err)
			})
		}
	}
	return nil
}
