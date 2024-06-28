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
		a := e.Group("a")
		{
			// サービス管理者向けAPI
			a.Use(w.APIMiddlewareAuthAuthe)
			photoStudios := a.Group("photo_studios")
			{
				photoStudios.POST("", w.APIPostPhotoStudios)
				photoStudio := photoStudios.Group(":photoStudioID")
				photoStudio.Use(w.APIMiddlewarePhotoStudio)
				{
					members := photoStudio.Group("members")
					{
						members.POST("", w.APIPostPhotoStudioMembers)
					}
					customers := photoStudio.Group("customers")
					{
						customers.GET("search", w.APIGetPhotoStudioCustomers)
					}
				}
			}
		}
	}

	{
		b := e.Group("b")
		// Authを担うAPI
		b.POST("login", w.AuthPostLogin)
	}

	{
		c := e.Group("c")
		// スーパーAPI
		c.POST("init", w.SuperPostInit)
	}
}
