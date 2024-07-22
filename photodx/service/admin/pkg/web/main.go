package web

import (
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/gateway/line/messaging"
	internal_infra_repository "github.com/suzuito/sandbox2-go/photodx/service/admin/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/usecase"
	auth_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/auth/pkg/businesslogic"
	authuser_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/authuser/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
	"gorm.io/gorm"
)

func MiddlewareAccessTokenAuthe(
	l *slog.Logger,
	u usecase.Usecase,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
		common_web.CtxSetAdminPrincipalAccessToken(ctx, dto.Principal)
		ctx.Next()
	}
}

func Main(
	e *gin.Engine,
	l *slog.Logger,
	db *gorm.DB,
	adminAccessTokenVerifier auth.JWTVerifier,
	authUserBusinessLogic authuser_businesslogic.ExposedBusinessLogic,
	authBusinessLogic auth_businesslogic.ExposedBusinessLogic,
	skipVerifyLINEWebhook bool,
) error {
	r := internal_infra_repository.Impl{
		GormDB:  db,
		NowFunc: time.Now,
	}
	u := usecase.Impl{
		BusinessLogic: &businesslogic.Impl{
			LINEMessagingAPIClient: &messaging.Impl{
				Cli: http.DefaultClient,
			},
			Repository: &r,
			L:          l,
		},
		CommonBusinessLogic: common_businesslogic.NewBusinessLogic(
			adminAccessTokenVerifier,
			nil,
		),
		AuthUserBusinessLogic: authUserBusinessLogic,
		AuthBusinessLogic:     authBusinessLogic,
		L:                     l,
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
	// スタジオ管理画面向けAPI
	admin := e.Group("admin")
	{
		admin.Use(MiddlewareAccessTokenAuthe(l, &u))
		{
			photoStudios := admin.Group("photo_studios")
			{
				photoStudio := photoStudios.Group(":photoStudioID")
				photoStudio.Use(
					common_web.MiddlewareAdminAccessTokenAutho(
						l,
						`
							permissions.exists(
    							p,
			                    p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "read".matches(p.action)
		                    ) && pathParams["photoStudioID"] == adminPrincipalPhotoStudioId
							`,
						&p,
					),
				)
				{
					lineLink := photoStudio.Group("line_link")
					lineLink.PUT(
						"activate",
						common_web.MiddlewareAdminAccessTokenAutho(
							l,
							`
								permissions.exists(
									p,
									p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "update".matches(p.action)
								)
								`,
							&p,
						),
						func(ctx *gin.Context) {
							principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
							dto, err := u.APIPutLINELinkActivate(ctx, principal)
							res(ctx, dto, err)
						},
					)
					lineLink.PUT(
						"deactivate",
						common_web.MiddlewareAdminAccessTokenAutho(
							l,
							`
								permissions.exists(
									p,
									p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "update".matches(p.action)
								)
								`,
							&p,
						),
						func(ctx *gin.Context) {
							principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
							dto, err := u.APIPutLINELinkDeactivate(ctx, principal)
							res(ctx, dto, err)
						},
					)
					lineLink.GET(
						"",
						common_web.MiddlewareAdminAccessTokenAutho(
							l,
							`
								permissions.exists(
									p,
									p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "read".matches(p.action)
								)
								`,
							&p,
						),
						func(ctx *gin.Context) {
							principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
							dto, err := u.APIGetLINELink(ctx, principal)
							res(ctx, dto, err)
						},
					)
					lineLink.PUT(
						"",
						common_web.MiddlewareAdminAccessTokenAutho(
							l,
							`
								permissions.exists(
									p,
									p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "update".matches(p.action)
								)
								`,
							&p,
						),
						func(ctx *gin.Context) {
							message := repository.SetLineLinkInfoArgument{}
							if err := ctx.BindJSON(&message); err != nil {
								p.JSON(
									ctx,
									http.StatusBadRequest,
									common_web.ResponseError{
										Message: err.Error(),
									},
								)
								return
							}
							principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
							dto, err := u.APIPutLINELink(ctx, principal, &message)
							res(ctx, dto, err)
						},
					)
				}
				{
					users := photoStudio.Group("users")
					users.GET(
						"",
						common_web.MiddlewareAdminAccessTokenAutho(
							l,
							`
								permissions.exists(
									p,
									p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "update".matches(p.action)
								)
								`,
							&p,
						),
						func(ctx *gin.Context) {
							query := struct {
								Offset int `form:"offset"`
							}{}
							if err := ctx.BindQuery(&query); err != nil {
								p.JSON(
									ctx,
									http.StatusBadRequest,
									common_web.ResponseError{
										Message: err.Error(),
									},
								)
								return
							}
							principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
							dto, err := u.APIGetPhotoStudioUsers(ctx, principal, query.Offset)
							res(ctx, dto, err)
						},
					)
					{
						user := users.Group(":userID")
						user.GET(
							"",
							common_web.MiddlewareAdminAccessTokenAutho(
								l,
								`
									permissions.exists(
										p,
										p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "update".matches(p.action)
									)
									`,
								&p,
							),
							func(ctx *gin.Context) {
								principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
								dto, err := u.APIGetPhotoStudioUser(ctx, principal, entity.UserID(ctx.Param("userID")))
								res(ctx, dto, err)
							},
						)
					}
				}
			}
		}
	}
	// Webhooks
	{
		wh := admin.Group("wh")
		wh.POST(
			"line_messaging_api_webhook/:photoStudioID",
			func(ctx *gin.Context) {
				photoStudioID := entity.PhotoStudioID(ctx.Param("photoStudioID"))
				body, err := io.ReadAll(ctx.Request.Body)
				if err != nil {
					p.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
						Message: "%+v",
					})
					return
				}
				err = u.APIPostLineMessagingAPIWebhook(
					ctx,
					photoStudioID,
					body,
					ctx.GetHeader("x-line-signature"),
					skipVerifyLINEWebhook,
				)
				res(ctx, struct{}{}, err)
			},
		)
	}

	// TODO 後で消す
	admin.POST("super_init", func(ctx *gin.Context) {
		dto, err := u.APIPostSuperInit(ctx)
		res(ctx, dto, err)
	})
	return nil
}
