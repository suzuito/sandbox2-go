package web

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetRouter(
	e *gin.Engine,
	w *Impl,
) {
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

	api := e.Group("api")
	api.Use(w.MiddlewareAuth0Authe)
	{
		api.GET("hoge", func(ctx *gin.Context) {
			claims := ctxGetAuth0ValidatedClaims(ctx)
			fmt.Println(claims)
			customClaims := claims.CustomClaims.(*Auth0CustomClaims)
			fmt.Println(customClaims)
		})
		photoStudios := api.Group("photo_studios")
		{
			photoStudio := photoStudios.Group(":photoStudioID")
			photoStudio.Use(w.MiddlewarePhotoStudio)
			{
				customers := photoStudio.Group("customers")
				{
					customers.GET("search", w.GetPhotoStudioCustomers)
				}
			}
		}
	}
}
