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
	e.Static("js", "blog2/internal/web/_js")
	e.Static("css", "blog2/internal/web/_css")
	e.Static("images", "blog2/internal/web/_images")
	e.Static("wasm", "blog2/internal/web/_wasm")
	e.GET("health", w.GetHealth)
	e.Use(w.MiddlewareAdminAuthe)
	e.NoRoute(w.NoRoute)
	e.GET("", w.GetTop)
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
	{
		gAdmin := e.Group("admin")
		gAdmin.Use(w.MiddlewareAdminAutho)
		gAdmin.GET("", w.GetAdminTop)
		{
			gAdminArticles := gAdmin.Group("articles")
			gAdminArticles.GET("", w.GetAdminArticles)
			gAdminArticles.POST("", w.PostAdminArticles)
			{
				gAdminArticle := gAdminArticles.Group(":articleID")
				gAdminArticle.Use(w.MiddlewareGetArticle)
				gAdminArticle.GET("", w.GetAdminArticle)
				gAdminArticle.PUT("", w.PutAdminArticle)
				gAdminArticle.PUT("markdown", w.PutAdminArticleMarkdown)
				gAdminArticle.POST("publish", w.PostAdminArticlePublish)
				gAdminArticle.DELETE("publish", w.DeleteAdminArticlePublish)
				gAdminArticle.POST("edit-tags", w.PostAdminArticleEditTags)
				{
					gAdminArticleTags := gAdminArticle.Group("tags")
					gAdminArticleTags.GET("", w.GetAdminArticleTags)
				}
			}
		}
	}
}
