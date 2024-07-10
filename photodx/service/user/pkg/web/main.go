package web

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
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
) error {
	u := usecase.Impl{
		BusinessLogic: &businesslogic.Impl{},
		CommonBusinessLogic: common_businesslogic.NewBusinessLogic(
			nil,
			userAccessTokenJWTVerifier,
		),
		L: l,
	}
	w := internal_web.Impl{
		U: &u,
		L: l,
		P: &presenter.Impl{},
	}
	app := e.Group("app")
	app.Use(w.MiddlewareAccessTokenAuthe)
	app.GET(
		"init",
		common_web.MiddlewareAccessTokenAutho(`
			permissions.exists(
				p,
				p.resource == "PhotoStudio" && "read".matches(p.action)
			)
		`, w.P),
		w.GetInit,
	)
	return nil
}
