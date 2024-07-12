package web

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/businesslogic"
	internal_infra_repository "github.com/suzuito/sandbox2-go/photodx/service/admin/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/admin/internal/web"
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
) error {
	repository := internal_infra_repository.Impl{
		GormDB:  db,
		NowFunc: time.Now,
	}
	u := usecase.Impl{
		BusinessLogic: &businesslogic.Impl{
			Repository: &repository,
			L:          l,
		},
		CommonBusinessLogic: common_businesslogic.NewBusinessLogic(
			adminAccessTokenVerifier,
			nil,
		),
		L: l,
	}
	w := internal_web.Impl{
		U: &u,
		L: l,
		P: &presenter.Impl{},
	}
	// スタジオ管理画面向けAPI
	admin := e.Group("admin")
	{
		admin.Use(w.MiddlewareAccessTokenAuthe)
		admin.GET(
			"init",
			common_web.MiddlewareAdminAccessTokenAutho(
				`
					permissions.exists(
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
			w.APIGetInit,
		)
		{
			photoStudios := admin.Group("photo_studios")
			{
				photoStudio := photoStudios.Group(":photoStudioID")
				photoStudio.Use(
					common_web.MiddlewareAdminAccessTokenAutho(
						`
							permissions.exists(
    							p,
			                    p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "read".matches(p.action)
		                    )
							`,
						w.P,
					),
					w.APIMiddlewarePhotoStudio,
				)
				{
					lineLink := photoStudio.Group("line_link")
					lineLink.POST(
						"",
						common_web.MiddlewareAdminAccessTokenAutho(
							`
								permissions.exists(
									p,
									p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "update".matches(p.action)
								)
								`,
							w.P,
						),
						w.APIPostLINELink,
					)
					lineLink.GET(
						"",
						common_web.MiddlewareAdminAccessTokenAutho(
							`
								permissions.exists(
									p,
									p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "read".matches(p.action)
								)
								`,
							w.P,
						),
						w.APIGetLINELink,
					)
					lineLink.DELETE(
						"",
						common_web.MiddlewareAdminAccessTokenAutho(
							`
								permissions.exists(
									p,
									p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "update".matches(p.action)
								)
								`,
							w.P,
						),
						w.APIDeleteLINELink,
					)
					lineLink.PUT(
						"messaging_api_channel_secret",
						common_web.MiddlewareAdminAccessTokenAutho(
							`
								permissions.exists(
									p,
									p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "update".matches(p.action)
								)
								`,
							w.P,
						),
						w.APIPutLINELinkMessagingAPIChannelSecret,
					)
				}
			}
		}
	}
	// Webhooks
	{
		wh := admin.Group("wh")
		wh.POST("line_messaging_api_webhook", w.APIPostLineMessagingAPIWebhook)
	}
	return nil
}
