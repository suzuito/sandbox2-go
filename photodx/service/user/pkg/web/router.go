package web

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
	"github.com/suzuito/sandbox2-go/photodx/service/user/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/user/internal/web"
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
	app := e.Group("app")
	app.Use(w.MiddlewareAccessTokenAuthe)
	app.GET(
		"init",
		w.MiddlewareAccessTokenAutho(`
			permissions.exists(
				p,
				p.resource == "PhotoStudio" && "read".matches(p.action)
			)
		`),
		w.GetInit,
	)
}
