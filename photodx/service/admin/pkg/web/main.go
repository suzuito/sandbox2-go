package web

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/usecase"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/admin/internal/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

func Main(
	e *gin.Engine,
	l *slog.Logger,
	adminAccessTokenJWTPublicKey string,
) error {
	adminAccessTokenJWTPublicKeyBytes, err := jwt.ParseRSAPublicKeyFromPEM([]byte(adminAccessTokenJWTPublicKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	b := businesslogic.Impl{
		L: l,
		AdminAccessTokenJWTVerifier: &auth.JWTVerifiers{
			Verifiers: []auth.JWTVerifier{
				&auth.JWTVerifierRS256{
					PublicKey: adminAccessTokenJWTPublicKeyBytes,
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
	return nil
}
