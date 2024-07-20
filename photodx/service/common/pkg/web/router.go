package web

import (
	"log/slog"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

func Main(
	e *gin.Engine,
	l *slog.Logger,
	corsAllowOrigins []string,
	corsAllowMethods []string,
	corsAllowHeaders []string,
	corsExposeHeaders []string,
) {
	var p presenter.Presenter = &presenter.Impl{}
	e.NoRoute(func(ctx *gin.Context) {
		p.JSON(ctx, http.StatusNotFound, ResponseError{
			Message: "not found",
		})
	})
	e.Use(func(ctx *gin.Context) {
		ctx.Header("X-Robots-Tag", "noindex")
		ctx.Next()
	})
	e.Use(cors.New(cors.Config{
		AllowOrigins:     corsAllowOrigins,
		AllowMethods:     corsAllowMethods,
		AllowHeaders:     corsAllowHeaders,
		ExposeHeaders:    corsExposeHeaders,
		AllowCredentials: true,
	}))
	e.GET("health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
}
