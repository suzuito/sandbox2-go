package web

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	admin_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/admin/pkg/businesslogic"
	auth_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/auth/pkg/businesslogic"
	authuser_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/authuser/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
	"github.com/suzuito/sandbox2-go/photodx/service/user/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/user/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/user/internal/web"
)

func Main(
	e *gin.Engine,
	l *slog.Logger,
	userAccessTokenJWTVerifier auth.JWTVerifier,
	authBusinessLogic auth_businesslogic.ExposedBusinessLogic,
	authUserBusinessLogic authuser_businesslogic.ExposedBusinessLogic,
	adminBusinessLogic admin_businesslogic.ExposedBusinessLogic,
) error {
	var u usecase.Usecase = &usecase.Impl{
		NowFunc:       time.Now,
		BusinessLogic: &businesslogic.Impl{},
		CommonBusinessLogic: common_businesslogic.NewBusinessLogic(
			nil,
			userAccessTokenJWTVerifier,
		),
		AuthBusinessLogic:     authBusinessLogic,
		AuthUserBusinessLogic: authUserBusinessLogic,
		AdminBusinessLogic:    adminBusinessLogic,
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
	app := e.Group("app")
	app.Use(func(ctx *gin.Context) {
		authorizationHeaderString := ctx.GetHeader("Authorization")
		authorizationHeaderParts := strings.Fields(authorizationHeaderString)
		if len(authorizationHeaderParts) != 2 || strings.ToLower(authorizationHeaderParts[0]) != "bearer" {
			ctx.Next()
			return
		}
		accessToken := authorizationHeaderParts[1]
		dto, err := u.MiddlewareAccessTokenAuthe(ctx, accessToken)
		if err != nil {
			l.Warn("accessToken authe is failed", "err", err)
			ctx.Next()
			return
		}
		common_web.CtxSetUserPrincipalAccessToken(ctx, dto.UserPrincipal)
		ctx.Next()
	})
	{
		photoStudios := app.Group("photo_studios")
		{
			photoStudio := photoStudios.Group(":photoStudioID")
			photoStudio.Use(func(ctx *gin.Context) {
				photoStudioID := common_entity.PhotoStudioID(ctx.Param("photoStudioID"))
				dto, err := u.MiddlewarePhotoStudio(ctx, photoStudioID)
				if err != nil {
					res(ctx, nil, err)
					ctx.Abort()
					return
				}
				internal_web.CtxSetPhotoStudio(ctx, dto.PhotoStudio)
				ctx.Next()
			})
			photoStudio.GET("", func(ctx *gin.Context) {
				res(ctx, internal_web.CtxGetPhotoStudio(ctx), nil)
			})
			{
				chatMessages := photoStudio.Group("messages")
				chatMessages.GET(
					"older",
					func(ctx *gin.Context) {
						photoStudioID := common_entity.PhotoStudioID(ctx.Param("photoStudioID"))
						userID := common_web.CtxGetUserPrincipalAccessToken(ctx).GetUserID()
						query := struct {
							Offset int `form:"offset"`
						}{}
						if err := ctx.BindQuery(&query); err != nil {
							p.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
								Message: err.Error(),
							})
							return
						}
						if query.Offset < 0 {
							query.Offset = 0
						}
						dto, err := u.APIGetOlderPhotoStudioChatMessages(
							ctx,
							photoStudioID,
							userID,
							query.Offset,
						)
						res(ctx, dto, err)
					},
				)
				chatMessages.GET(
					"older_by_id",
					func(ctx *gin.Context) {
						photoStudioID := common_entity.PhotoStudioID(ctx.Param("photoStudioID"))
						userID := common_web.CtxGetUserPrincipalAccessToken(ctx).GetUserID()
						query := struct {
							ID common_entity.ChatMessageID `form:"id"`
						}{}
						if err := ctx.BindQuery(&query); err != nil {
							p.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
								Message: err.Error(),
							})
							return
						}
						if query.ID == "" {
							p.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
								Message: "id is empty",
							})
							return
						}
						dto, err := u.APIGetOlderPhotoStudioChatMessagesByID(
							ctx,
							photoStudioID,
							userID,
							query.ID,
						)
						res(ctx, dto, err)
					},
				)
				chatMessages.POST(
					"",
					common_web.MiddlewareUserAccessTokenAutho(
						l,
						`
						permissions.exists(
							p,
							p.resource == "PhotoStudio" && "read".matches(p.action)
						)
					`,
						&p,
					),
					func(ctx *gin.Context) {
						photoStudioID := common_entity.PhotoStudioID(ctx.Param("photoStudioID"))
						message := usecase.InputAPIPostPhotoStudioMessages{}
						if err := ctx.BindJSON(&message); err != nil {
							p.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
								Message: err.Error(),
							})
							return
						}
						_, skipPushMessage := ctx.GetQuery("skipPushMessage")
						dto, err := u.APIPostPhotoStudioMessages(
							ctx,
							common_web.CtxGetUserPrincipalAccessToken(ctx),
							photoStudioID,
							&message,
							skipPushMessage,
						)
						res(ctx, dto, err)
					},
				)
			}
		}
	}
	app.GET(
		"init",
		common_web.MiddlewareUserAccessTokenAutho(
			l,
			`
			permissions.exists(
				p,
				p.resource == "PhotoStudio" && "read".matches(p.action)
			)
		`, &p),
	)
	return nil
}
