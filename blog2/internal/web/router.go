package web

import (
	"path"

	"github.com/gin-gonic/gin"
)

func SetRouter(
	e *gin.Engine,
	w *Impl,
) {
	e.LoadHTMLGlob(path.Join("blog2/internal/web/_templates", "*"))
	e.Static("css", "blog2/internal/web/_css")
	e.Static("images", "blog2/internal/web/_images")
	e.GET("health", w.GetHealth)
	e.GET("", func(ctx *gin.Context) {
	})
	{
		gArticles := e.Group("articles")
		gArticles.GET("", func(ctx *gin.Context) {})
		{
			gArticle := gArticles.Group(":articleID")
			gArticle.GET("", func(ctx *gin.Context) {})
			gArticle.GET("edit", func(ctx *gin.Context) {})
			gArticle.PUT("edit", func(ctx *gin.Context) {})
		}
	}
	{
		gTags := e.Group("tags")
		gTags.GET("", func(ctx *gin.Context) {})
		{
			gTag := gTags.Group(":tagID")
			gTag.GET("", func(ctx *gin.Context) {})
			gTag.GET("edit", func(ctx *gin.Context) {})
		}
	}
	e.GET("login", w.GetLogin)
	e.POST("login", func(ctx *gin.Context) {})
	e.POST("logout", func(ctx *gin.Context) {})
}
