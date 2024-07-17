package web

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/businesslogic"
	internal_infra_repository "github.com/suzuito/sandbox2-go/photodx/service/admin/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/admin/internal/web"
	authuser_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/authuser/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
	"gorm.io/gorm"
)

func Main(
	e *gin.Engine,
	l *slog.Logger,
	db *gorm.DB,
	adminAccessTokenVerifier auth.JWTVerifier,
	authUserBusinessLogic authuser_businesslogic.ExposedBusinessLogic,
) error {
	r := internal_infra_repository.Impl{
		GormDB:  db,
		NowFunc: time.Now,
	}
	u := usecase.Impl{
		BusinessLogic: &businesslogic.Impl{
			Repository: &r,
			L:          l,
		},
		CommonBusinessLogic: common_businesslogic.NewBusinessLogic(
			adminAccessTokenVerifier,
			nil,
		),
		AuthUserBusinessLogic: authUserBusinessLogic,
		L:                     l,
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
	// スタジオ管理画面向けAPI
	admin := e.Group("admin")
	{
		admin.Use(w.MiddlewareAccessTokenAuthe)
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
						w.P,
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
							w.P,
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
							w.P,
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
							w.P,
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
							w.P,
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
			}
		}
	}
	// Webhooks
	{
		wh := admin.Group("wh")
		wh.POST("line_messaging_api_webhook/:photoStudioID", w.APIPostLineMessagingAPIWebhook)
	}
	return nil
}
