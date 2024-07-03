package web

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetRouter(
	e *gin.Engine,
	w *Impl,
) {
	e.NoRoute(func(ctx *gin.Context) {
		w.P.JSON(ctx, http.StatusNotFound, ResponseError{
			Message: "not found",
		})
	})
	e.Use(func(ctx *gin.Context) {
		ctx.Header("X-Robots-Tag", "noindex")
		ctx.Next()
	})
	e.Use(cors.New(cors.Config{
		AllowOrigins:     w.CorsAllowOrigins,
		AllowMethods:     w.CorsAllowMethods,
		AllowHeaders:     w.CorsAllowHeaders,
		ExposeHeaders:    w.CorsExposeHeaders,
		AllowCredentials: true,
	}))

	e.GET("health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	{
		// ユーザー画面向けAPI
		x := e.Group("x")
		{
			x.Use(func(ctx *gin.Context) {
			}) // auth0 authe
			x.GET(
				"init",
				w.MiddlewareAccessTokenAutho(
					`
					permissions.exists(
						p,
						p.resource == "PhotoStudio" && "read".matches(p.action)
					)
					`,
				),
				func(ctx *gin.Context) {},
			)
		}
	}

	{
		// スタジオ管理画面向けAPI
		a := e.Group("a")
		{
			a.Use(w.MiddlewareAccessTokenAuthe)
			a.GET(
				"init",
				w.MiddlewareAccessTokenAutho(
					`
					permissions.exists(
						p,
						p.resource == "PhotoStudio" && principalPhotoStudioId.matches(p.target) && "read".matches(p.action)
					) &&
					permissions.exists(
						p,
						p.resource == "PhotoStudioMember" && principalPhotoStudioMemberId.matches(p.target) && "read".matches(p.action)
					)
					`,
				),
				w.APIGetInit,
			)
			{
				photoStudios := a.Group("photo_studios")
				// photoStudios.POST("", w.APIPostPhotoStudios)
				{
					photoStudio := photoStudios.Group(":photoStudioID")
					photoStudio.Use(
						w.MiddlewareAccessTokenAutho(
							`
							permissions.exists(
    							p,
			                    p.resource == "PhotoStudio" && principalPhotoStudioId.matches(p.target) && "read".matches(p.action)
		                    )
							`,
						),
						w.APIMiddlewarePhotoStudio,
					)
				}
			}
		}
	}

	{
		// Authを担うAPI
		b := e.Group("b")
		b.POST("login", w.AuthPostLogin)
		{
			x := b.Group("x")
			x.Use(w.MiddlewareRefreshTokenAuthe)
			x.Use(w.MiddlewareRefreshTokenAutho)
			x.POST(
				"refresh",
				w.AuthPostRefresh,
			)
		}
	}

	{
		// スーパーAPI
		c := e.Group("c")
		c.POST("init", w.SuperPostInit)
	}
}
