package web

import (
	"log/slog"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/web/internal/presenter"
	"github.com/suzuito/sandbox2-go/blog2/pkg/usecase"
)

type Impl struct {
	U                    usecase.Usecase
	P                    presenter.Presenter
	L                    *slog.Logger
	NoIndex              bool
	AdminToken           string
	BaseURLFile          string
	BaseURLFileThumbnail string
	SiteOrigin           string
	GoogleTagManagerID   string
}

func NewPresenter() presenter.Presenter {
	return &presenter.Impl{}
}

func SetRouter(
	e *gin.Engine,
	w *Impl,
) {
	if w.NoIndex {
		e.Use(func(ctx *gin.Context) {
			ctx.Header("X-Robots-Tag", "noindex")
			ctx.Next()
		})
	}
	e.LoadHTMLGlob(path.Join("blog2/internal/web", "*.html"))
	e.Static("js", "blog2/internal/web/_js")
	e.Static("css", "blog2/internal/web/_css")
	e.Static("images", "blog2/internal/web/_images")
	e.GET("health", w.PageHealth)
	e.Use(w.MiddlewareAdminAuthe)
	e.NoRoute(w.PageNoRoute)
	e.GET("", w.PageTop)
	{
		gArticles := e.Group("articles")
		gArticles.GET("", w.PageArticles)
		{
			gArticle := gArticles.Group(":articleID")
			gArticle.Use(w.GetMiddlewareGetArticle(true))
			gArticle.GET("", w.PageArticle)
		}
	}
	{
		gTags := e.Group("tags")
		gTags.GET("", func(ctx *gin.Context) {})
		{
			gTag := gTags.Group(":tagID")
			gTag.GET("", func(ctx *gin.Context) {})
		}
	}
	{
		gAdmin := e.Group("admin")
		gAdmin.Use(w.MiddlewareAdminAutho)
		gAdmin.GET("", w.PageAdminTop)
		{
			gAdminArticles := gAdmin.Group("articles")
			gAdminArticles.GET("", w.PageAdminArticles)
			gAdminArticles.POST("", w.PostAdminArticles)
			{
				gAdminArticle := gAdminArticles.Group(":articleID")
				gAdminArticle.Use(w.GetMiddlewareGetArticle(false))
				gAdminArticle.GET("", w.PageAdminArticle)
			}
		}
	}

	gAPI := e.Group("api")
	{
		gAdmin := gAPI.Group("admin")
		gAdmin.Use(w.MiddlewareAdminAutho)
		{
			gAdminArticles := gAdmin.Group("articles")
			{
				gAdminArticle := gAdminArticles.Group(":articleID")
				{
					gAdminArticle.Use(w.GetMiddlewareGetArticle(false))
					gAdminArticle.PUT("", w.APIPutAdminArticle)
					gAdminArticle.POST("/edit-tags", w.APIPostAdminArticleEditTags)
					gAdminArticle.PUT("/markdown", w.APIPutAdminArticleMarkdown)
				}
			}
			{
				gAdminFiles := gAdmin.Group("files")
				gAdminFiles.POST("", w.APIPostAdminFiles)
				gAdminFiles.GET("", w.APIGetAdminFiles)
			}
		}
	}
}
