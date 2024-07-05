package web

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/admin/internal/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

func SetRouter(
	e *gin.Engine,
	l *slog.Logger,
	b businesslogic.BusinessLogic,
) {
	u := usecase.Impl{
		B: b,
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
			w.MiddlewareAccessTokenAutho(
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
			),
			w.APIGetInit,
		)
		{
			photoStudios := admin.Group("photo_studios")
			// photoStudios.POST("", w.APIPostPhotoStudios)
			{
				photoStudio := photoStudios.Group(":photoStudioID")
				photoStudio.Use(
					w.MiddlewareAccessTokenAutho(
						`
							permissions.exists(
    							p,
			                    p.resource == "PhotoStudio" && adminPrincipalPhotoStudioId.matches(p.target) && "read".matches(p.action)
		                    )
							`,
					),
					w.APIMiddlewarePhotoStudio,
				)
			}
		}
	}
}
