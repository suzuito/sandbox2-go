package web

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/admin/internal/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

func Main(
	e *gin.Engine,
	l *slog.Logger,
	adminAccessTokenVerifier auth.JWTVerifier,
) error {
	u := usecase.Impl{
		BusinessLogic: &businesslogic.Impl{
			L: l,
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
			// photoStudios.POST("", w.APIPostPhotoStudios)
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
			}
		}
	}
	return nil
}
