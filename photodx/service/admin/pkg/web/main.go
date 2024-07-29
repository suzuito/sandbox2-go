package web

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/gateway/line/messaging"
	internal_infra_repository "github.com/suzuito/sandbox2-go/photodx/service/admin/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/usecase"
	auth_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/auth/pkg/businesslogic"
	authuser_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/authuser/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
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
) error {
	r := internal_infra_repository.Impl{
		GormDB:  db,
		NowFunc: time.Now,
	}
	var u usecase.Usecase = &usecase.Impl{
		NowFunc: time.Now,
		BusinessLogic: &businesslogic.Impl{
			LINEMessagingAPIClient: &messaging.Impl{
				Cli: http.DefaultClient,
			},
			Repository:            &r,
			GenerateChatMessageID: &proc.IDGeneratorImpl{},
			L:                     l,
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
		admin.Use(MiddlewareAccessTokenAuthe(l, u))
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
				{
					chats := photoStudio.Group("chats")
					chats.GET(
						"",
						func(ctx *gin.Context) {
							query := struct {
								Offset int `form:"query"`
							}{}
							if err := ctx.BindQuery(&query); err != nil {
								p.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{Message: err.Error()})
								return
							}
							dto, err := u.APIGetPhotoStudioChats(
								ctx,
								common_web.CtxGetAdminPrincipalAccessToken(ctx).GetPhotoStudioID(),
								query.Offset,
							)
							res(ctx, dto, err)
						},
					)
					{
						chat := chats.Group(":userID")
						chat.GET(
							"",
							func(ctx *gin.Context) {
								dto, err := u.APIGetPhotoStudioChat(
									ctx,
									common_web.CtxGetAdminPrincipalAccessToken(ctx).GetPhotoStudioID(),
									entity.UserID(ctx.Param("userID")),
								)
								res(ctx, dto, err)
							},
						)
						{
							chatMessages := chat.Group("messages")
							chatMessages.GET(
								"older",
								func(ctx *gin.Context) {
									userID := entity.UserID(ctx.Param("userID"))
									photoStudioID := common_web.CtxGetAdminPrincipalAccessToken(ctx).GetPhotoStudioID()
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
									userID := entity.UserID(ctx.Param("userID"))
									photoStudioID := common_web.CtxGetAdminPrincipalAccessToken(ctx).GetPhotoStudioID()
									query := struct {
										ID entity.ChatMessageID `form:"id"`
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
								func(ctx *gin.Context) {
									userID := entity.UserID(ctx.Param("userID"))
									principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
									message := struct {
										Text string `json:"text"`
									}{}
									if err := ctx.BindJSON(&message); err != nil {
										p.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
											Message: err.Error(),
										})
										return
									}
									_, skipPushMessage := ctx.GetQuery("skipPushMessage")
									dto, err := u.APIPostPhotoStudioChatMessages(
										ctx,
										principal.GetPhotoStudioID(),
										userID,
										principal.GetPhotoStudioMemberID(),
										message.Text,
										skipPushMessage,
									)
									res(ctx, dto, err)
								},
							)
						}
					}
				}
			}
		}
	}

	// TODO 後で消す
	admin.POST("super_init", func(ctx *gin.Context) {
		dto, err := u.APIPostSuperInit(ctx)
		res(ctx, dto, err)
	})
	return nil
}
