package web

import (
	"log/slog"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/oauth2loginflow"
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
		OAuth2ProviderLINE: &oauth2loginflow.Provider{
			ClientID:     "2005761043",
			ClientSecret: "3250327d6ab0c0f92938d37e6ff87750",
		},
	}
	oauth2RedirectURL, _ := url.Parse("http://localhost:8082/authuser/x/callback")
	w := internal_web.Impl{
		U:                 &u,
		P:                 &presenter.Impl{},
		OAuth2RedirectURL: *oauth2RedirectURL,
	}
	authuser := e.Group("authuser")
	// authuser.POST("login", func(ctx *gin.Context) {}) // Password login
	{
		x := authuser.Group("x")
		x.GET("callback", w.GetCallback)
	}
	{
		authorize := authuser.Group("authorize")
		authorize.GET("line", w.GetAuthorizeLine)
	}
}
