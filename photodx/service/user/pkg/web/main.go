package web

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
	"github.com/suzuito/sandbox2-go/photodx/service/user/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/user/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/user/internal/web"
)

func Main(
	e *gin.Engine,
	l *slog.Logger,
	userAccessTokenJWTPublicKey string,
) error {
	userAccessTokenJWTPublicKeyBytes, err := jwt.ParseRSAPublicKeyFromPEM([]byte(userAccessTokenJWTPublicKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	b := businesslogic.Impl{
		UserAccessTokenJWTVerifier: &auth.JWTVerifiers{
			Verifiers: []auth.JWTVerifier{
				&auth.JWTVerifierRS256{
					PublicKey: userAccessTokenJWTPublicKeyBytes,
				},
			},
		},
	}
	u := usecase.Impl{
		B: &b,
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
	return nil
}
