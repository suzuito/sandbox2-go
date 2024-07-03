package web

import (
	"log/slog"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	internal_web "github.com/suzuito/sandbox2-go/photodx/service/common/internal/web"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

func SetRouter(
	e *gin.Engine,
	l *slog.Logger,
	corsAllowOrigins []string,
	corsAllowMethods []string,
	corsAllowHeaders []string,
	corsExposeHeaders []string,
) {
	w := internal_web.Impl{
		L:                 l,
		P:                 &presenter.Impl{},
		CorsAllowOrigins:  corsAllowOrigins,
		CorsAllowMethods:  corsAllowMethods,
		CorsAllowHeaders:  corsAllowHeaders,
		CorsExposeHeaders: corsExposeHeaders,
	}
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
}
